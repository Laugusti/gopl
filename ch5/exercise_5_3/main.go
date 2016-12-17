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
	printText(doc)
}

// visit appends to links each link found in n and returns the result.
func printText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}
	if n.Data != "script" && n.Data != "style" && n.FirstChild != nil {
		printText(n.FirstChild)
	}
	if n.NextSibling != nil {
		printText(n.NextSibling)
	}
}
