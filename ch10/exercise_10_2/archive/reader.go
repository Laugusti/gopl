package archive

import (
	"io"
	"os"
)

type MultiFileReader interface {
	Next() (*File, error)
}

type File struct {
	Name   string
	Reader io.Reader
}

// Creates a decompressing reader
func NewReader(file *os.File) (MultiFileReader, error) {
	f, err := getFormat(file)
	if err != nil {
		return nil, err
	}
	return f.decompressor(file)
}
