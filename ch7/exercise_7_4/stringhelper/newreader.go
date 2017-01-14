package stringhelper

import (
	"io"
)

type reader struct {
	bytes []byte
	pos   int
}

func (r *reader) Read(p []byte) (n int, err error) {
	n = copy(p, r.bytes[r.pos:])
	r.pos += n
	if r.pos >= len(r.bytes) {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	r := &reader{bytes: []byte(s)}
	return r
}
