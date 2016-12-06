// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	categoryCounts := make(map[string]int) // counts of Unicode categories
	var utflen [utf8.UTFMax + 1]int        // counts of lengths of UTF-8 encodings
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			categoryCounts["letter"]++
		case unicode.IsDigit(r):
			categoryCounts["digit"]++
		case unicode.IsSpace(r):
			categoryCounts["space"]++
		case unicode.IsPunct(r):
			categoryCounts["punctuation"]++
		case unicode.IsControl(r):
			categoryCounts["control"]++
		case unicode.IsGraphic(r):
			categoryCounts["graphic"]++
		case unicode.IsMark(r):
			categoryCounts["mark"]++
		default:
			categoryCounts["other"]++
		}
		utflen[n]++
	}
	fmt.Printf("category\tcount\n")
	for c, n := range categoryCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalide UTF-8 characters\n", invalid)
	}
}
