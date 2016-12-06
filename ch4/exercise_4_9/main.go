// WordCount computes counts of words.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCount := make(map[string]int) // counts of each word in an input text

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		wordCount[in.Text()]++
	}
	fmt.Printf("word\tcount\n")
	for w, n := range wordCount {
		fmt.Printf("%q\t%d\n", w, n)
	}
}
