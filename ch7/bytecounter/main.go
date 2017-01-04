package main

import (
	"fmt"

	"github.com/Laugusti/gopl/bytecounter/bytecounter"
)

func main() {
	var c bytecounter.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
