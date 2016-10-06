// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of name files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		if hasDuplicates(counts) {
			fmt.Println("os.Stdin")
		}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			if hasDuplicates(counts) {
				fmt.Println(arg)
			}
			f.Close()
			counts = make(map[string]int)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func hasDuplicates(counts map[string]int) bool {
	for _, count := range counts {
		if count > 1 {
			return true
		}
	}
	return false
}
