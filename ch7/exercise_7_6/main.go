package main

import (
	"flag"
	"fmt"

	"github.com/Laugusti/gopl/ch7/exercise_7_6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(temp)
}
