// comma inserts commas in a floating-point number string with optional sign.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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
	prefix := ""
	if strings.HasPrefix(s, "+") {
		s = s[1:]
		prefix = "+"
	} else if strings.HasPrefix(s, "-") {
		s = s[1:]
		prefix = "-"
	}
	suffix := ""
	if index := strings.Index(s, "."); index >= 0 {
		suffix = s[index:]
		s = s[:index]
	}
	for i, v := range s {
		buf.WriteRune(v)
		if len(s) > 3 && i != len(s)-1 && len(s[i+1:])%3 == 0 {
			buf.WriteByte(',')
		}
	}
	return prefix + buf.String() + suffix
}
