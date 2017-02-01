// Package ftpdirectory provides for functions for a directory in a ftp server.
package ftpdirectory

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const root = "/tmp"

type Dir string

// PrintWorkingDirectory returns the current working directory
func (d *Dir) PrintWorkingDirectory() string {
	return filepath.Join(root, string(*d))
}

// ChangeDirectory changes the current working directory
func (d *Dir) ChangeDirectory(p string) error {
	dir := getRootedPath(string(*d), p)

	// return an error if the argument is not a valid
	// directory
	fi, err := os.Stat(dir)
	if err != nil || !fi.IsDir() {
		return fmt.Errorf("not a valid directory")
	}

	// change the current working directory
	*d = Dir(filepath.Clean(dir[len(root):]))
	return nil
}

// ListFiles performs operation analogous to the unix ls command
func (d *Dir) ListFiles(p string) ([]string, error) {
	path := getRootedPath(string(*d), p)
	fi, err := os.Stat(path)

	// returns the *PathError (if any)
	if err != nil {
		return nil, err
	}

	if fi.IsDir() {
		// argument is a directory, return contents
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		return f.Readdirnames(-1)
	} else {
		// argument is a file, return filename
		return []string{filepath.Base(path)}, nil
	}
}

// GetFile returns the bytes of the file, or an error
func (d *Dir) GetFile(fpath string) ([]byte, error) {
	path := getRootedPath(string(*d), fpath)
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("%q is a directory", fpath)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// StoreFile creates the file fpath and copies writes the contents
// from the provided reader
func (d *Dir) StoreFile(fpath string, r io.Reader) error {
	path := getRootedPath(string(*d), fpath)
	fi, err := os.Stat(filepath.Dir(path))
	if err != nil || !fi.IsDir() {
		return fmt.Errorf("not a valid directory")
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, r)
	return err
}

// getRootedPath returns the true path using root
func getRootedPath(wd, s string) string {
	// use relative to cwd, if not absolute
	if filepath.Clean(s)[0] != filepath.Separator {
		s = filepath.Join(wd, s)
	}

	// return path rooted at root
	rootedDir := filepath.Join(string(filepath.Separator), s)
	rootedDir = filepath.Join(root, rootedDir)
	return rootedDir
}
