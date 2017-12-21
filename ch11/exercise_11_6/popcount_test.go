package popcount

import (
	"testing"

	"github.com/Laugusti/gopl/ch2/exercise_2_4"
	"github.com/Laugusti/gopl/ch2/exercise_2_5"
	"github.com/Laugusti/gopl/ch2/popcount"
)

func benchmark(b *testing.B, impl func(uint64) int) {
	for i := 0; i < b.N; i++ {
		impl(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchmark(b, popcount.PopCount)
}

func BenchmarkPopCountByShifting(b *testing.B) {
	benchmark(b, popcountshift.PopCountShift)
}

func BenchmarkPopCountByClearing(b *testing.B) {
	benchmark(b, popcountclear.PopCountClear)
}
