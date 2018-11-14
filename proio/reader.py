import gzip
import io
import lz4.frame
import struct
import sys

import google.protobuf.descriptor_pool as descriptor_pool

from .event import Event
import proio.proto as proto
from .writer import magic_bytes

class Reader(object):
    """
    Reader for proio files

    This class can be used with the `with` statement, and it also may be used
    as an iterator that sequentially iterates all events.  A filename may be
    omitted in favor of specifying `fileobj`.

    :param string filename: name of input file to read
    :param fileobj: file object to read from

    :example:

    .. code-block:: python

        with proio.Reader('input.proio') as reader:
            for event in reader:
                ...

    """

    def __init__(self, filename = None, fileobj = None):
        if filename is None:
            if fileobj is not None:
                self._stream_reader = fileobj
            else:
                self._stream_reader = io.BytesIO(b'')
        else:
            self._stream_reader = open(filename, 'rb')
            self._close_file = True

        self._bucket_index = 0
        self._bucket_header = None
        self._bucket_reader = None
        self._metadata = {}

    def __enter__(self):
        return self

    def __exit__(self, exception_type, exception_value, traceback):
        self.close()

    def __iter__(self):
        return self

    def __next__(self):
        """
        :return: the next event
        :rtype: Event
        """
        event = None

        # use skip() to ensure that we land on a non-empty bucket
        self.skip(0)
        if self._bucket_header is not None:
            if self._bucket_reader is None:
                self._read_bucket()
            event = self._read_from_bucket()

        if event is None:
            raise StopIteration
        return event

    if sys.version_info[0] == 2:
        def next(self):
            return self.__next__()

    def close(self):
        """
        closes the underlying input file object.
        """
        try:
            if self._close_file:
                self._stream_reader.close()
        except AttributeError:
            pass

    def skip(self, n_events):
        """
        skips the next `n_events` events.

        :param int n_events: number of events to skip
        :return: number of events skipped
        :rtype: int
        """
        n_skipped = 0

        start_index = self._bucket_index
        self._bucket_index += n_events
        while self._bucket_header is None or self._bucket_index >= self._bucket_header.nEvents:
            if self._bucket_header is not None:
                n_bucket_events = self._bucket_header.nEvents
                self._bucket_index -= n_bucket_events
                n_skipped += n_bucket_events - start_index

                # skip the bucket bytes on the stream if they haven't been read
                # into memory already
                if n_bucket_events > 0 and self._bucket_reader is None:
                    try:
                        self._stream_reader.seek(self._bucket_header.bucketSize, 1)
                    except OSError:
                        self._stream_reader.read(self._bucket_header.bucketSize)
            self._read_header()
            if self._bucket_header is None:
                return n_skipped
            start_index = 0
        n_skipped += self._bucket_index - start_index

        return n_skipped

    def seek_to_start(self):
        """
        seeks, if possible, to the start of the input file object.  This can be
        used along with :func:`skip` to directly access events.

        :return: success
        :rtype: boolean
        """
        if self._stream_reader.seekable():
            self._stream_reader.seek(0, 0)
            self._metadata = {}
            self._bucket_index = 0
            self._read_header()
            return True
        return False

    def _read_header(self):
        self._bucket_evts_read = 0
        self._bucket_reader = None
        self._bucket_header = None
        
        n = self._sync_to_magic()
        if n < len(magic_bytes):
            return

        header_size = struct.unpack("I", self._stream_reader.read(4))[0]
        header_string = self._stream_reader.read(header_size)
        if len(header_string) != header_size:
            return
        self._bucket_header = proto.BucketHeader.FromString(header_string)

        # set metadata for future events
        for key, value in self._bucket_header.metadata.items():
            self._metadata[key] = value

        # add descriptors to pool
        for fd_bytes in self._bucket_header.fileDescriptor:
            try:
                descriptor_pool.Default().AddSerializedFile(fd_bytes)
            except TypeError:
                # ignore cases where types were already defined by another
                # FileDescriptorProto
                pass

    def _read_bucket(self):
        bucket = self._stream_reader.read(self._bucket_header.bucketSize)

        if len(bucket) != self._bucket_header.bucketSize:
            raise EOFError

        if self._bucket_header.compression == proto.BucketHeader.GZIP:
            self._bucket_reader = gzip.GzipFile(fileobj = io.BytesIO(bucket), mode = 'rb')
        elif self._bucket_header.compression == proto.BucketHeader.LZ4:
            try:
                uncomp_bytes, _ = lz4.frame.decompress(bucket)
            except ValueError:
                uncomp_bytes = lz4.frame.decompress(bucket)
            self._bucket_reader = io.BytesIO(uncomp_bytes)
        elif self._bucket_header.compression == proto.BucketHeader.NONE:
            self._bucket_reader = io.BytesIO(bucket)
        else:
            raise UnknownBucketCompError

    def _read_from_bucket(self):
        event = None

        while self._bucket_evts_read <= self._bucket_index:
            proto_size_buf = self._bucket_reader.read(4)
            if len(proto_size_buf) != 4:
                raise CorruptBucketError('Unable to read event size')

            proto_size = struct.unpack("I", proto_size_buf)[0]
            proto_buf = self._bucket_reader.read(proto_size)
            if len(proto_buf) != proto_size:
                raise CorruptBucketError('Unable to read event data')

            if self._bucket_evts_read == self._bucket_index:
                event_proto = proto.Event.FromString(proto_buf)
                event = Event(proto_obj = event_proto)
                for key, value in self._metadata.items():
                    event.metadata[key] = value

            self._bucket_evts_read += 1
        self._bucket_index += 1

        return event

    def _sync_to_magic(self):
        n_read = 0
        while True:
            magic_byte = self._stream_reader.read(1)
            if len(magic_byte) != 1:
                return -1
            n_read += 1

            if magic_byte == magic_bytes[0]:
                goodSeq = True
                for i in range(1, len(magic_bytes)):
                    magic_byte = self._stream_reader.read(1)
                    if len(magic_byte) != 1:
                        return -1
                    n_read += 1

                    if magic_byte != magic_bytes[i]:
                        goodSeq = False
                        break
                if goodSeq:
                    break

        return n_read

class UnknownBucketCompError(Exception):
    """
    raised when a bucket is compressed with an uknown type
    """
    pass

class CorruptBucketError(Exception):
    """
    raised when there is a problem reading from a bucket
    """
    pass
