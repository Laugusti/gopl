package zip

import (
	"archive/zip"
	"io"
	"os"

	"github.com/Laugusti/gopl/ch10/exercise_10_2/archive"
)

// reader satisfies archive.MultiFileReader
type reader struct {
	zipReader   *zip.Reader
	currentFile int
}

func (r *reader) Next() (*archive.File, error) {
	r.currentFile++
	if r.currentFile >= len(r.zipReader.File) {
		return nil, io.EOF
	}
	f := r.zipReader.File[r.currentFile]
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	return &archive.File{f.Name, rc}, nil
}

// Create a zip reader from the File parameter
func NewReader(f *os.File) (archive.MultiFileReader, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	zr, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return nil, err
	}
	return &reader{zr, -1}, nil
}

// register the zip decompressor
func init() {
	archive.RegisterFormat(archive.Magic{"PK", 0}, NewReader)
}
