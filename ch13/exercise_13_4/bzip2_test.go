package bzip

import (
	"bytes"
	"compress/bzip2"
	"io/ioutil"
	"strings"
	"sync"
	"testing"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	count    = 1000
)

func TestConcurrentWrites(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewWriter(buf)
	var n sync.WaitGroup
	for i := 0; i < count; i++ {
		n.Add(1)
		go func(i int) {
			if _, err := w.Write([]byte(alphabet)); err != nil {
				t.Fatal(err)
			}
			n.Done()
		}(i)
	}
	n.Wait()
	b, err := ioutil.ReadAll(bzip2.NewReader(buf))
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != len(alphabet)*count {
		t.Errorf("expected %d uncompressed bytes written, got %d", len(alphabet)*count, len(b))
	}
	if string(b) != strings.Repeat(alphabet, count) {
		t.Errorf("unexpected data written")
	}
}
