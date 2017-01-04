package main

import (
	"fmt"

	"github.com/Laugusti/gopl/ch6/exercise_6_1/intset"
)

func main() {
	var x intset.IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)

	fmt.Println(&x)       // "{1 2 3}"
	fmt.Println(x.Copy()) // "{1 2 3}"
	fmt.Println(x.Len())  // "3"
	x.Remove(2)
	fmt.Println(&x)      // "{1 3}"
	fmt.Println(x.Len()) // "2"
	x.Clear()
	fmt.Println(&x) // "{}"
}
