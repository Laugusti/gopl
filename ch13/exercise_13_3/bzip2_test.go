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
			if i < count/2 {
				w.Write([]byte(alphabet))
			} else {
				w.Close()
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
	if len(b)%len(alphabet) != 0 {
		t.Errorf("expected multiple of %d uncompressed bytes written, got %d", len(alphabet), len(b))
	}
	if string(b) != strings.Repeat(alphabet, len(b)/len(alphabet)) {
		t.Errorf("unexpected data written")
	}
	// t.Log(string(b))
}
