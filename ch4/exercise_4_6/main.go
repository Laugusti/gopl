// joinSpaces squashes (in-place) each run of adjacent Unicode spaces in a UTF-8-encoded []byte slice into a single ASCII space.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	data := []byte("This  is  a  string!")
	fmt.Printf("%q\n", joinSpaces(data)) // "This is a string!"
	fmt.Printf("%q\n", data)             // "This is a string!ng!"
}

func joinSpaces(bytes []byte) []byte {
	length := 0
	lastRuneIsSpace := false
	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		if unicode.IsSpace(r) {
			if !lastRuneIsSpace {
				lastRuneIsSpace = true
				bytes[length] = ' '
				length++
			}
		} else {
			lastRuneIsSpace = false
			for j := 0; j < size; j++ {
				bytes[length+j] = bytes[i+j]
			}
			length += size
		}
		i += size
	}
	return bytes[:length]
}
