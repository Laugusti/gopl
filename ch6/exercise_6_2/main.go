package main

import (
	"fmt"

	"github.com/Laugusti/gopl/ch6/exercise_6_2/intset"
)

func main() {
	var x intset.IntSet
	x.AddAll(1, 2, 11000, 3)
	fmt.Println(&x) // "{1 2 3 11000}"
}
