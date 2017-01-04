package countingwriter

import (
	"testing"

	"github.com/Laugusti/gopl/ch7/bytecounter/bytecounter"
)

func getCountIgnoreError(n int, err error) int64 {
	return int64(n)
}
func TestCountingWriter(t *testing.T) {
	var total int64
	var w bytecounter.ByteCounter
	cw, bytesWritten := CountingWriter(&w)
	if *bytesWritten != int64(w) || *bytesWritten != total {
		t.Errorf("bytecounter: %d, countingwriter: %d, expected: %d", w, *bytesWritten, 0)
	}

	total += getCountIgnoreError(cw.Write([]byte("123456789")))
	if *bytesWritten != int64(w) || *bytesWritten != total {
		t.Errorf("bytecounter: %d, countingwriter: %d, expected: %d", w, *bytesWritten, total)
	}

	total += getCountIgnoreError(cw.Write([]byte("123456789")))
	if *bytesWritten != int64(w) || *bytesWritten != total {
		t.Errorf("bytecounter: %d, countingwriter: %d, expected: %d", w, *bytesWritten, total)
	}

	total += getCountIgnoreError(cw.Write([]byte("")))
	if *bytesWritten != int64(w) || *bytesWritten != total {
		t.Errorf("bytecounter: %d, countingwriter: %d, expected: %d", w, *bytesWritten, total)
	}
}
