package main

import (
	"flag"
	"fmt"

	"github.com/Laugusti/gopl/ch7/tempflag/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(temp)
}
