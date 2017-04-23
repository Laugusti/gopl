package tar

import (
	"archive/tar"
	"os"

	"github.com/Laugusti/gopl/ch10/exercise_10_2/archive"
)

// reader satisfies archive.MultiFileReader
type reader struct {
	tarReader *tar.Reader
}

func (r *reader) Next() (*archive.File, error) {
	header, err := r.tarReader.Next()
	if err != nil {
		return nil, err
	}
	return &archive.File{header.Name, r.tarReader}, nil
}

// Create a tar reader from the File parameter
func NewReader(f *os.File) (archive.MultiFileReader, error) {
	return &reader{tar.NewReader(f)}, nil
}

// register the tar decompressor
func init() {
	archive.RegisterFormat(archive.Magic{"ustar", 257}, NewReader)
}
