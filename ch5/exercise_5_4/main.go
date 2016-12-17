package main

import "fmt"
import "golang.org/x/net/html"
import "os"

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if isLinkNode(n) {
		for _, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func isLinkNode(n *html.Node) bool {
	if n.Type != html.ElementNode {
		return false
	}
	return n.Data == "a" || n.Data == "link" || n.Data == "script" || n.Data == "img"
}
