package countingwriter

import "io"

type Writer struct {
	inner io.Writer
	count *int64
}

func (w Writer) Write(p []byte) (int, error) {
	n, err := w.inner.Write(p)
	*w.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var x int64
	cw := Writer{w, &x}
	return cw, cw.count
}
