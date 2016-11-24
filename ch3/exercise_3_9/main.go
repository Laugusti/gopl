// Creates a web server that renders a PNG image of the Newton's fractal for p(z) = z^4 - 1.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strconv"
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
	http.HandleFunc("/", drawFractal)
	http.ListenAndServe(":8000", nil)
}

func drawFractal(w http.ResponseWriter, r *http.Request) {
	x_min := getParameter("xmin", r, xmin)
	x_max := getParameter("xmax", r, xmax)
	y_min := getParameter("ymin", r, ymin)
	y_max := getParameter("ymax", r, ymax)
	pWidth := int(getParameter("width", r, width))
	pHeight := int(getParameter("height", r, height))

	img := image.NewRGBA(image.Rect(0, 0, pWidth, pHeight))

	for px := 0; px < pWidth; px++ {
		for py := 0; py < pHeight; py++ {
			x := float64(px)/float64(pWidth)*(x_max-x_min) + x_min
			y := float64(py)/float64(pHeight)*(y_max-y_min) + y_min

			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newtonsFractal(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func getParameter(parameter string, r *http.Request, defaultValue float64) float64 {
	if val := r.URL.Query().Get(parameter); len(val) != 0 {
		result, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return defaultValue
		}
		return result
	}
	return defaultValue
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
