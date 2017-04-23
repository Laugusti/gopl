package gzip

import (
	"compress/gzip"
	"encoding/hex"
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

// Create a zip reader from the File parameter
func NewReader(f *os.File) (archive.MultiFileReader, error) {
	zr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	name := strings.TrimRight(filepath.Base(f.Name()), ".gz")
	return &reader{zr, name, false}, nil
}

func init() {
	b, err := hex.DecodeString("1f8b")
	if err != nil {
		panic(err)
	}
	archive.RegisterFormat(archive.Magic{string(b), 0}, NewReader)
}
