// Mandel brot emits a PNG image of the MandelBrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	iterations             = 200
	contrast               = 15
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img.Set(px, py, superSample(px, py))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func getColor(iteration uint8) color.Color {
	return color.RGBA{contrast * iteration, contrast * iteration, 255, 255}
}

func mandelbrot(z complex128) color.Color {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return getColor(n)
		}
	}
	return color.Black
}

func getSample(px int, py int, sample int) complex128 {
	x := float64(px*2+sample/2)/(2*width)*(xmax-xmin) + xmin
	y := float64(py*2+sample%2)/(2*height)*(ymax-ymin) + ymin
	return complex(x, y)
}

func averageColor(color1 color.Color, color2 color.Color, color3 color.Color, color4 color.Color) color.Color {
	r1, g1, b1, _ := color1.RGBA()
	r2, g2, b2, _ := color2.RGBA()
	r3, g3, b3, _ := color3.RGBA()
	r4, g4, b4, _ := color4.RGBA()

	r := uint8((int(r1) + int(r2) + int(r3) + int(r4)) / (0x101 * 4))
	g := uint8((int(g1) + int(g2) + int(g3) + int(g4)) / (0x101 * 4))
	b := uint8((int(b1) + int(b2) + int(b3) + int(b4)) / (0x101 * 4))
	return color.RGBA{r, g, b, 255}
}

func superSample(px int, py int) color.Color {
	color1 := mandelbrot(getSample(px, py, 0))
	color2 := mandelbrot(getSample(px, py, 1))
	color3 := mandelbrot(getSample(px, py, 2))
	color4 := mandelbrot(getSample(px, py, 3))

	return averageColor(color1, color2, color3, color4)
}
