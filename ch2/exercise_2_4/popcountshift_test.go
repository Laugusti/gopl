package popcountshift

import (
	"testing"

	"github.com/Laugusti/gopl/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(0x1234567890ABCDEF)
	}
}

func TestPopCountShift(t *testing.T) {
	var input uint64 = 0x1234567890ABCDEF
	if PopCountShift(input) != popcount.PopCount(input) {
		t.Errorf("PopCountShift != PopCount for input %q", input)
	}
}
