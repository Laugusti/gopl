// reverse reverses a slice of bytes in place
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := []byte("Hello, 世界")
	reverse(a)
	fmt.Printf("%q\n", a) // "界世 ,olleH"
}

func reverse(bytes []byte) {
	var firstSize, lastSize int
	for i, j := 0, len(bytes); utf8.RuneCount(bytes[i:j]) > 1; i, j = i+firstSize, j-lastSize {
		firstRune, size := utf8.DecodeRune(bytes[i:])
		firstSize = size
		lastRune, size := utf8.DecodeLastRune(bytes[:j])
		lastSize = size
		if firstSize != lastSize {
			copy(bytes[i+lastSize:], bytes[i+firstSize:j-lastSize])
		}
		copy(bytes[i:], []byte(string(lastRune)))
		copy(bytes[j-firstSize:], []byte(string(firstRune)))
		firstSize, lastSize = lastSize, firstSize
	}
}
