package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var red = color.RGBA{0xFF, 0x0, 0x0, 0xFF}
var yellow = color.RGBA{0xFF, 0xFF, 0x0, 0xFF}
var green = color.RGBA{0x0, 0xFF, 0x0, 0xFF}
var cyan = color.RGBA{0x0, 0xFF, 0xFF, 0xFF}
var blue = color.RGBA{0x0, 0x0, 0xFF, 0xFF}

var palette = []color.Color{color.Black, red, yellow, green, cyan, blue}

func main() {
	http.HandleFunc("/", lissajous)
	http.ListenAndServe("raspberrypi2:8000", nil)
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	cycles := 5   // numbr of complete x oscillator revolutions
	res := 0.001  // angular resolution
	size := 100   // image canvas convers [-size..+size]
	nframes := 64 // number of animation frames
	delay := 8    // delay between frames in 10ms units

	if val := r.URL.Query().Get("cycles"); len(val) != 0 {
		parsedVal, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			fmt.Fprintf(w, "Invalid cycles parameter: %v\n", val)
		}
		cycles = int(parsedVal)
	}

	if val := r.URL.Query().Get("res"); len(val) != 0 {
		parsedVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Fprintf(w, "Invalid res parameter: %v\n", val)
		}
		res = parsedVal
	}

	if val := r.URL.Query().Get("size"); len(val) != 0 {
		parsedVal, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			fmt.Fprintf(w, "Invalid size parameter: %v\n", val)
		}
		size = int(parsedVal)
	}

	if val := r.URL.Query().Get("nframes"); len(val) != 0 {
		parsedVal, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			fmt.Fprintf(w, "Invalid nframes parameter: %v\n", val)
		}
		nframes = int(parsedVal)
	}

	if val := r.URL.Query().Get("delay"); len(val) != 0 {
		parsedVal, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			fmt.Fprintf(w, "Invalid delay parameter: %v\n", val)
		}
		delay = int(parsedVal)
	}

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := uint8(rand.Intn(5) + 1)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim) // NOTE: ignoring encoding errors
}
