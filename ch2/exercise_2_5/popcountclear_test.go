package popcountclear

import (
	"testing"

	"github.com/Laugusti/gopl/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(0x1234567890ABCDEF)
	}
}

func TestPopCountClear(t *testing.T) {
	var input uint64 = 0x1234567890ABCDEF
	if PopCountClear(input) != popcount.PopCount(input) {
		t.Errorf("PopCountClear != PopCount for input %q", input)
	}
}
