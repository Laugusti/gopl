package stringhelper

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	r := LimitReader(strings.NewReader("abc"), 2)
	buf := make([]byte, 1024)

	n, err := r.Read(buf)
	if n != 2 || err != io.EOF {
		t.Fail()
	}
}
