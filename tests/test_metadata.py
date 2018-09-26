import io
import pytest

import proio

def test_header1():
    buf = io.BytesIO(b'')
    with proio.Writer(fileobj = buf) as writer:
        event = proio.Event()
        writer.push(event)
        writer.push(event)
        writer.flush()
        writer.push(event)

    buf.seek(0, 0)

    with proio.Reader(fileobj = buf) as reader:
        reader.skip(0)
        assert(reader._bucket_header.nEvents == 2)
        reader.skip(2)
        assert(reader._bucket_header.nEvents == 1)
        reader.skip(1)
        assert(reader._bucket_header is None)

def test_push_update1():
    buf = io.BytesIO(b'')
    with proio.Writer(fileobj = buf) as writer:
        writer.push_metadata('key1', b'value1')
        writer.push_metadata('key2', b'value2')
        event = proio.Event()
        writer.push(event)
        writer.push_metadata('key2', b'value3')
        writer.push(event)
        writer.push_metadata('key1', b'value4')
        writer.push_metadata('key2', b'value5')
        writer.push(event)

    buf.seek(0, 0)
    with proio.Reader(fileobj = buf) as reader:
        event1 = next(reader)
        event2 = next(reader)
        event3 = next(reader)
        assert(event1.metadata['key1'] == b'value1')
        assert(event1.metadata['key2'] == b'value2')
        assert(event2.metadata['key1'] == b'value1')
        assert(event2.metadata['key2'] == b'value3')
        assert(event3.metadata['key1'] == b'value4')
        assert(event3.metadata['key2'] == b'value5')

    buf = io.BytesIO(b'')
    with proio.Writer(fileobj = buf) as writer:
        writer.push(event1)
        writer.push(event2)
        writer.push(event3)

    buf.seek(0, 0)
    with proio.Reader(fileobj = buf) as reader:
        event1 = next(reader)
        event2 = next(reader)
        event3 = next(reader)
        assert(event1.metadata['key1'] == b'value1')
        assert(event1.metadata['key2'] == b'value2')
        assert(event2.metadata['key1'] == b'value1')
        assert(event2.metadata['key2'] == b'value3')
        assert(event3.metadata['key1'] == b'value4')
        assert(event3.metadata['key2'] == b'value5')

