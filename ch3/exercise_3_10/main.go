// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: comma NUMBER")
		os.Exit(1)
	}
	fmt.Println(comma(os.Args[1]))
}

func comma(s string) string {
	var buf bytes.Buffer
	for i, v := range s {
		buf.WriteRune(v)
		if len(s) > 3 && i != len(s)-1 && len(s[i+1:])%3 == 0 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
