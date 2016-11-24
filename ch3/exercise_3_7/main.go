// Emits a PNG image of the Newton's fractal for p(z) = z^4 - 1.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	iterations             = 200
	approx_zero            = 0.0000001
)

var roots [4]complex128
var foundRoots uint8

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for px := 0; px < width; px++ {
		for py := 0; py < height; py++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			y := float64(py)/height*(ymax-ymin) + ymin

			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newtonsFractal(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
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

func newtonsFractal(z complex128) color.Color {
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
