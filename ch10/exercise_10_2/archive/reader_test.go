package archive_test

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Laugusti/gopl/ch10/exercise_10_2/archive"
	_ "github.com/Laugusti/gopl/ch10/exercise_10_2/archive/bzip2"
	_ "github.com/Laugusti/gopl/ch10/exercise_10_2/archive/gzip"
	_ "github.com/Laugusti/gopl/ch10/exercise_10_2/archive/tar"
	_ "github.com/Laugusti/gopl/ch10/exercise_10_2/archive/zip"
)

type content struct {
	filename, content string
}

var archiveTests = []struct {
	filename         string
	expectedContents []content
}{
	{"./testdata/file.tar", []content{content{"file.dat", "this is a test"}}},
	{"./testdata/file.zip", []content{content{"file.dat", "this is a test"}}},
	{"./testdata/file.dat.bz2", []content{content{"file.dat", "this is a test"}}},
	{"./testdata/file.dat.gz", []content{content{"file.dat", "this is a test"}}},
}

func TestReader(t *testing.T) {
	for _, ar := range archiveTests {
		f, err := os.Open(ar.filename)
		if err != nil {
			t.Error(err)
		}
		r, err := archive.NewReader(f)
		if err != nil {
			t.Errorf("creating reader for %s: %v", ar.filename, err)
		}
		for _, c := range ar.expectedContents {
			af, err := r.Next()
			if err != nil {
				t.Error(err)
			}
			if c.filename != af.Name {
				t.Errorf("expected file: %q, actual file %q", c.filename, af.Name)
			}
			b, err := ioutil.ReadAll(af.Reader)
			if err != nil {
				t.Error(err)
			}
			if c.content != string(b) {
				t.Errorf("expected content: %q, actual content: %q", c.content, string(b))
			}
		}
		if _, err = r.Next(); err != io.EOF {
			t.Errorf("expected io.EOF, got %v", err)
		}
		f.Close()
	}
}
