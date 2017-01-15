package stringhelper

import (
	"io"
)

type reader struct {
	innerReader io.Reader
	limit       int
	pos         int
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.innerReader.Read(p[:r.pos+r.limit])
	r.pos += n
	if r.pos >= r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int) io.Reader {
	limitReader := &reader{r, n, 0}
	return limitReader
}
