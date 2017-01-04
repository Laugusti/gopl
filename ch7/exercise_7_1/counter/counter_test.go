package counter

import (
	"fmt"
	"testing"
)

var testStr = `one two three
four five sixy
seven eight nine`

func TestByteCounter(t *testing.T) {
	var c ByteCounter
	fmt.Fprintf(&c, "")
	if c != 0 {
		t.Errorf("ByteCounter failed: expected %d, got %d", 0, c)
	}

	fmt.Fprintf(&c, testStr)
	if c != 45 {
		t.Errorf("ByteCounter failed: expected %d, got %d", 45, c)
	}

	fmt.Fprintf(&c, testStr)
	if c != 90 {
		t.Errorf("ByteCounter failed: expected %d, got %d", 90, c)
	}
}

func TestWordCounter(t *testing.T) {
	var c WordCounter
	fmt.Fprintf(&c, "")
	if c != 0 {
		t.Errorf("WordCounter failed: expected %d, got %d", 0, c)
	}

	fmt.Fprintf(&c, testStr)
	if c != 9 {
		t.Errorf("WordCounter failed: expected %d, got %d", 9, c)
	}

	fmt.Fprintf(&c, testStr)
	if c != 18 {
		t.Errorf("WordCounter failed: expected %d, got %d", 18, c)
	}
}

func TestLineCounter(t *testing.T) {
	var c LineCounter
	if c != 0 {
		t.Errorf("LineCounter failed: expected %d, got %d", 0, c)
	}

	fmt.Fprintf(&c, testStr)
	if c != 3 {
		t.Errorf("LineCounter failed: expected %d, got %d", 3, c)
	}

	fmt.Fprintf(&c, testStr)
	if c != 6 {
		t.Errorf("LineCounter failed: expected %d, got %d", 6, c)
	}
}
