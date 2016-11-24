// Emits a PNG image of the Newton's fractal for p(z) = z^4 - 1.
package complex64Fractal

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	iterations             = 200
	approx_zero            = 0.0000001
)

var roots [4]complex64
var foundRoots uint8

func NewtonsFractal(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			y := float32(py)/height*(ymax-ymin) + ymin

			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newtonsMethod(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func newtonsMethod(z complex64) color.Color {
	var p complex64
	for n := uint8(0); n < iterations; n++ {
		// p(z) = z^4 - 1
		p = Pow64(z, 4) - 1
		// p ~= 0
		if Abs64(p) <= approx_zero {
			return getColor(z)
		} else {
			// Newton's method
			z = z - (Pow64(z, 4)-1)/(4*Pow64(z, 3))
		}
	}
	return color.Black
}

func getColor(z complex64) color.Color {
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 0, 255, 255}
	blue := color.RGBA{0, 255, 0, 255}
	cyan := color.RGBA{0, 255, 255, 255}
	colors := []color.Color{red, green, blue, cyan}

	switch foundRoots {
	case 4:
		if Abs64(roots[3]-z) <= approx_zero {
			return colors[3]
		}
		fallthrough
	case 3:
		if Abs64(roots[2]-z) <= approx_zero {
			return colors[2]
		}
		fallthrough
	case 2:
		if Abs64(roots[1]-z) <= approx_zero {
			return colors[1]
		}
		fallthrough
	case 1:
		if Abs64(roots[0]-z) <= approx_zero {
			return colors[0]
		}
		fallthrough
	default:
		roots[foundRoots] = z
		foundRoots++
		return colors[foundRoots-1]
	}
}

func Pow64(x, y complex64) complex64 {
	return complex64(cmplx.Pow(complex128(x), complex128(y)))
}

func Abs64(x complex64) float32 {
	return float32(cmplx.Abs(complex128(x)))
}
