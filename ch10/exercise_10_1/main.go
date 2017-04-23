// Extennd the jpeg program so that it converts any supported input format
// to any output format. using image.Decode to detect the input format and
// a flag to select the output format.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var oFormat = flag.String("output", "jpg", "output image format")

type encodeFunc func(io.Writer, image.Image) error

// map of formats and their encoding functions
var formats = map[string]encodeFunc{
	"jpg": func(w io.Writer, m image.Image) error { return jpeg.Encode(w, m, &jpeg.Options{Quality: 95}) },
	"png": png.Encode,
	"gif": func(w io.Writer, m image.Image) error { return gif.Encode(w, m, &gif.Options{NumColors: 256}) },
}

func main() {
	flag.Parse()
	if formats[*oFormat] == nil {
		fmt.Fprintln(os.Stderr, "Unsupported output format")
		os.Exit(1)
	}
	if err := convertImage(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
	}
}

func convertImage(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return formats[*oFormat](out, img)
}
