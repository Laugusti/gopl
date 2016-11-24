package fractal_test

import (
	"io/ioutil"
	"testing"

	"github.com/Laugusti/gopl/ch3/exercise_3_8/bigFloatFractal"
	"github.com/Laugusti/gopl/ch3/exercise_3_8/bigRatFractal"
	"github.com/Laugusti/gopl/ch3/exercise_3_8/complex128Fractal"
	"github.com/Laugusti/gopl/ch3/exercise_3_8/complex64Fractal"
)

func BenchmarkComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		complex64Fractal.NewtonsFractal(ioutil.Discard)
	}
}

func BenchmarkComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		complex128Fractal.NewtonsFractal(ioutil.Discard)
	}
}

func BenchmarkBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigFloatFractal.NewtonsFractal(ioutil.Discard)
	}
}

func BenchmarkBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bigRatFractal.NewtonsFractal(ioutil.Discard)
	}
}
