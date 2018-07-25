package bzip

import (
	"bytes"
	"compress/bzip2"
	"io/ioutil"
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
			w.Write([]byte(alphabet[i%26 : i%26+1]))
			n.Done()
		}(i)
	}
	n.Wait()
	w.Close()
	b, err := ioutil.ReadAll(bzip2.NewReader(buf))
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != count {
		t.Errorf("expected %d unencrypted bytes written, got %d", count, len(b))
	}
	// t.Log(string(b))
}
