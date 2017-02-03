package mandlebrot

import (
	"testing"
)

func BenchmarkDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawImage()
	}
}

func BenchmarkDrawN1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawImageN(1)
	}
}

func BenchmarkDrawN2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawImageN(2)
	}
}

func BenchmarkDrawN4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawImageN(4)
	}
}

func BenchmarkDrawN8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawImageN(8)
	}
}

func BenchmarkDrawN16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawImageN(16)
	}
}
