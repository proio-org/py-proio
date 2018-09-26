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
