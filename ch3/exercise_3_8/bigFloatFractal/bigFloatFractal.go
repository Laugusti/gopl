// Emits a PNG image of the Newton's fractal for p(z) = z^4 - 1.
package bigFloatFractal

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	iterations             = 200
	approx_zero            = 0.0000001
)

var roots [4]complex128
var foundRoots uint8

func NewtonsFractal(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			x, _ := mapToPlane(px, width, xmin, xmax).Float64()
			y, _ := mapToPlane(py, height, ymin, ymax).Float64()

			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newtonsMethod(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mapToPlane(p, units, min, max int) *big.Float {
	result := big.NewFloat(float64(p))
	result.Quo(result, big.NewFloat(float64(units)))
	result.Mul(result, big.NewFloat(float64(max-min)))
	result.Add(result, big.NewFloat(float64(min)))
	return result
}

func newtonsMethod(z complex128) color.Color {
	var p complex128
	for n := uint8(0); n < iterations; n++ {
		// p(z) = z^4 - 1
		p = cmplx.Pow(z, 4) - 1
		// p ~= 0
		if cmplx.Abs(p) <= approx_zero {
			return getColor(z)
		} else {
			// Newton's method
			z = z - (cmplx.Pow(z, 4)-1)/(4*cmplx.Pow(z, 3))
		}
	}
	return color.Black
}

func getColor(z complex128) color.Color {
	red := color.RGBA{255, 0, 0, 255}
	green := color.RGBA{0, 0, 255, 255}
	blue := color.RGBA{0, 255, 0, 255}
	cyan := color.RGBA{0, 255, 255, 255}
	colors := []color.Color{red, green, blue, cyan}

	switch foundRoots {
	case 4:
		if cmplx.Abs(roots[3]-z) <= approx_zero {
			return colors[3]
		}
		fallthrough
	case 3:
		if cmplx.Abs(roots[2]-z) <= approx_zero {
			return colors[2]
		}
		fallthrough
	case 2:
		if cmplx.Abs(roots[1]-z) <= approx_zero {
			return colors[1]
		}
		fallthrough
	case 1:
		if cmplx.Abs(roots[0]-z) <= approx_zero {
			return colors[0]
		}
		fallthrough
	default:
		roots[foundRoots] = z
		foundRoots++
		return colors[foundRoots-1]
	}
}
