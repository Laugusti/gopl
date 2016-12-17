package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for tag, count := range visit(nil, doc) {
		fmt.Println(tag, count)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(tagCount map[string]int, n *html.Node) map[string]int {
	if tagCount == nil {
		tagCount = make(map[string]int)
	}
	if n.Type == html.ElementNode {
		tagCount[n.Data]++
	}
	if n.FirstChild != nil {
		tagCount = visit(tagCount, n.FirstChild)
	}
	if n.NextSibling != nil {
		tagCount = visit(tagCount, n.NextSibling)
	}
	return tagCount
}
