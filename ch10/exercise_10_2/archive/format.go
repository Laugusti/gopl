// Define a generic archive file-reading function capable of reading ZIP files
// (archive/zip) and POSIX tar files (archive/tar). Use a registration mechanism
// similar to the one described above so that support for each file format can
// be plugged in using blank imports.
package archive

import (
	"errors"
	"os"
)

// unsupported format error
var ErrFormat = errors.New("archive: unknown format")

// Magic holds the magic string and the offset
type Magic struct {
	M      string
	Offset int
}

// determines if a reader matches Magic
func (m Magic) match(f *os.File) bool {
	// get magic from file
	b := make([]byte, len(m.M))
	_, err := f.ReadAt(b, int64(m.Offset))
	if err != nil {
		return false
	}
	// reset file position
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return false
	}
	// return true if magic matches
	return m.M == string(b)
}

// format contains the magic number of the file and it's decompressor
type format struct {
	magic        Magic
	decompressor Decompressor
}

// TODO: need to guard??
var formats = []format{}

// Decompressor creates a reader form the file parameter
type Decompressor func(*os.File) (MultiFileReader, error)

// RegisterFormat registers decompressor for the magic number
func RegisterFormat(m Magic, d Decompressor) {
	formats = append(formats, format{m, d})
}

// retrieves the corresponding format for the io.Reader
func getFormat(file *os.File) (format, error) {
	for _, f := range formats {
		if f.magic.match(file) {
			return f, nil
		}
	}
	return format{}, ErrFormat
}
