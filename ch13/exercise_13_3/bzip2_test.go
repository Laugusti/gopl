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
	count    = 10000
)

func TestConcurrentWrites(t *testing.T) {
	buf := new(bytes.Buffer)
	w := NewWriter(buf)
	var n sync.WaitGroup
	for i := 0; i < count; i++ {
		n.Add(1)
		go func(i int) {
			_, err := w.Write([]byte(alphabet))
			if err != nil {
				t.Fatal(err)
			}
			n.Done()
		}(i)
	}
	n.Wait()
	w.Close()
	b, err := ioutil.ReadAll(bzip2.NewReader(buf))
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != count*len(alphabet) {
		t.Errorf("expected %d uncompressed bytes written, got %d", count*len(alphabet), len(b))
	}
	if string(b) != strings.Repeat(alphabet, count) {
		t.Errorf("unexpected data written")
	}
	t.Log(string(b))
}
