// Package explorefs provides functions for listing files and subdirectories
package explorefs

import (
	"os"
	"path/filepath"
)

// GetSubDirectories returns a list of directories under the specified path
func GetSubDirectories(path string) ([]string, error) {
	var subdirs []string

	err := readdir(path, func(fi os.FileInfo) {
		if fi.IsDir() {
			subdirs = append(subdirs, filepath.Join(path, fi.Name()))
		}
	})
	if err != nil {
		return nil, err
	}
	return subdirs, nil
}

func GetFiles(path string) ([]string, error) {
	var files []string

	err := readdir(path, func(fi os.FileInfo) {
		files = append(files, fi.Name())
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func readdir(path string, exec func(os.FileInfo)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	listFI, err := file.Readdir(-1)
	if err != nil {
		return err
	}

	for _, fi := range listFI {
		exec(fi)
	}
	file.Close()
	return nil
}
