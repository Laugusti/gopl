// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var red = color.RGBA{0xFF, 0x0, 0x0, 0xFF}
var yellow = color.RGBA{0xFF, 0xFF, 0x0, 0xFF}
var green = color.RGBA{0x0, 0xFF, 0x0, 0xFF}
var cyan = color.RGBA{0x0, 0xFF, 0xFF, 0xFF}
var blue = color.RGBA{0x0, 0x0, 0xFF, 0xFF}

var palette = []color.Color{color.Black, red, yellow, green, cyan, blue}

const (
	blackIndex = 0 // first color in the palette
	greenIndex = 1 // next color in the palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // numbr of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas convers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(5)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
