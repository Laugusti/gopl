package popcountloop

import (
	"testing"

	"github.com/Laugusti/gopl/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1234567890ABCDEF)
	}
}

func TestPopCountLoop(t *testing.T) {
	var input uint64 = 0x1234567890ABCDEF
	if PopCountLoop(input) != popcount.PopCount(input) {
		t.Errorf("PopCountLoop != PopCount for input %q", input)
	}
}
