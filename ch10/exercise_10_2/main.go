package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Laugusti/gopl/ch10/exercise_10_2/archive"
	_ "github.com/Laugusti/gopl/ch10/exercise_10_2/archive/tar"
	_ "github.com/Laugusti/gopl/ch10/exercise_10_2/archive/zip"
)

func main() {
	f, err := os.Open("./archive/testdata/file.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mfr, err := archive.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	for {
		af, err := mfr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("File: %s\n", af.Name)
		// NOTE: ignoring errors
		io.Copy(os.Stdout, af.Reader)
	}
}
