// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"sync"
	"time"
)

type writer struct {
	w  io.Writer // underlying output stream
	mu sync.Mutex
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.Writer {
	return &writer{w: out}
}

func (w *writer) Write(data []byte) (int, error) {
	w.mu.Lock()
	defer func() { w.mu.Unlock() }()
	cmd := exec.Command("bzip2", "-z")
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stdout = w.w
	if err := cmd.Start(); err != nil {
		return 0, err
	}

	// goroutine to wait on command and send error over channel
	ch := make(chan error)
	go func() {
		err := cmd.Wait()
		ch <- err
	}()

	// wait for timeout or for command to complete
	select {
	case <-time.After(5 * time.Second):
		return 0, fmt.Errorf("command timeout")
	case err := <-ch:
		return 0, err
	}
}
