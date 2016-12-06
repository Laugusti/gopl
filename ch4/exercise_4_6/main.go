// joinSpaces squashes (in-place) each run of adjacent Unicode spaces in a UTF-8-encoded []byte slice into a single ASCII space.
package main

import (
	"fmt"
	"unicode"
)

func main() {
	data := []byte("This  is  a  string!")
	fmt.Printf("%q\n", joinSpaces(data)) // "This is a string!"
	fmt.Printf("%q\n", data)             // "This is a string!ng!"
}

func joinSpaces(bytes []byte) []byte {
	i := 0
	lastRuneIsSpace := false
	for _, b := range bytes {
		if unicode.IsSpace(rune(b)) {
			if !lastRuneIsSpace {
				lastRuneIsSpace = true
				bytes[i] = ' '
				i++
			}
		} else {
			lastRuneIsSpace = false
			bytes[i] = b
			i++
		}
	}
	return bytes[:i]
}
