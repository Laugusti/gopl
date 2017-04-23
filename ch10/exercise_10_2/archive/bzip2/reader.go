package bzip2

import (
	"compress/bzip2"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/Laugusti/gopl/ch10/exercise_10_2/archive"
)

// reader satisfies archive.MultiFileReader
type reader struct {
	r    io.Reader
	name string
	done bool
}

func (r *reader) Next() (*archive.File, error) {
	if r.done {
		return nil, io.EOF
	}
	r.done = true
	return &archive.File{r.name, r.r}, nil
}

// Create a bzip2 reader from the File parameter
func NewReader(f *os.File) (archive.MultiFileReader, error) {
	name := strings.TrimRight(filepath.Base(f.Name()), ".bz2")
	return &reader{bzip2.NewReader(f), name, false}, nil
}

// register the bzip2 decompressor
func init() {
	archive.RegisterFormat(archive.Magic{"BZ", 0}, NewReader)
}
