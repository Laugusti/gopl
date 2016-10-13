// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 600            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 2.0                 //axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange //pixels per x or y unit
	zscale        = xyscale / 2         // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.5' "+
		"width='%d' height='%d'>", width, height)
	min, max := getMinMax()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			number := ax + ay + bx + by + cx + cy + dx + dy
			if math.IsNaN(number) || math.IsInf(number, 1) || math.IsInf(number, -1) {
				continue
			}

			z := getZ(i, j)
			colorValue := (z - min) / (max - min)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=%q/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, getColor(colorValue))
		}
	}
	fmt.Println("</svg>")
}

func getColor(value float64) string {
	red := uint8(0)
	green := uint8(0)
	blue := uint8(0)

	if value > 0.5 {
		if value > .75 {
			red = 255
			green = uint8((1 - value) * 4 * 255)
		} else {
			green = 255
			red = uint8(math.Abs(0.5-value) * 4 * 255)
		}
	} else {
		if value > .25 {
			green = 255
			blue = uint8((0.5 - value) * 4 * 255)
		} else {
			blue = 255
			green = uint8(value * 4 * 255)
		}
	}

	return fmt.Sprintf("#%02x%02x%02x", red, green, blue)
}

func getMinMax() (float64, float64) {
	var min, max float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := getZ(i, j)
			if z > max {
				max = z
			}
			if z < min {
				min = z
			}
		}
	}
	return min, max
}

func getZ(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	return f(x, y)
}

func corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	//Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	//return .5*math.Cos(x/2) + math.Sin(y/4)
	return math.Pow(x, 3) - 3*x*y*y
}
