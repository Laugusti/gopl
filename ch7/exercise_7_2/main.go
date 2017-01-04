package main

import (
	"fmt"

	"github.com/Laugusti/gopl/ch7/exercise_7_1/counter"
	"github.com/Laugusti/gopl/ch7/exercise_7_2/countingwriter"
)

func main() {
	var wc counter.WordCounter
	cw, count := countingwriter.CountingWriter(&wc)

	fmt.Fprintf(cw, "Hello, 世界")
	fmt.Println(wc)     // "2"
	fmt.Println(*count) // "13"

	wc = 0 // reset word counter
	fmt.Fprintf(cw, "Hello, 世界")
	fmt.Println(wc)     // "2"
	fmt.Println(*count) // "26"
}
