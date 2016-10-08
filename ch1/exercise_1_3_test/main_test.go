package benchmark

import (
	"os"
	"strings"
	"testing"
)

func BenchmarkLoopAndConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")
	}
}
