package main

import (
	"fmt"

	"github.com/Laugusti/gopl/ch7/exercise_7_1/counter"
)

func main() {
	var bCounter counter.ByteCounter
	bCounter.Write([]byte("Hello,\n世界"))
	fmt.Println(bCounter) // "13"

	var wCounter counter.WordCounter
	wCounter.Write([]byte("Hello,\n世界"))
	fmt.Println(wCounter) // "2"

	var lCounter counter.LineCounter
	lCounter.Write([]byte("Hello,\n世界"))
	fmt.Println(lCounter) // "2"
}
