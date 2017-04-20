// Mandel brot emits a PNG image of the MandelBrot fractal.
package mandlebrot

import (
	"image"
	"image/color"
	"math/cmplx"
	"sync"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

// creates mandelbrot image using a single goroutine
func drawImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	return img
}

// creates mandelbrot image using n goroutines
func drawImageN(n int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	type pixel struct{ x, y int }
	pixelChan := make(chan pixel, width*height)
	go func() {
		for py := 0; py < height; py++ {
			for px := 0; px < width; px++ {
				pixelChan <- pixel{px, py}
			}
		}
		close(pixelChan)
	}()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for p := range pixelChan {
				x := float64(p.x)/width*(xmax-xmin) + xmin
				y := float64(p.y)/height*(ymax-ymin) + ymin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(p.x, p.y, mandelbrot(z))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
