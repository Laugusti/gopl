// Package htmldom provides functions to search an HTML document.
package htmldom

import "golang.org/x/net/html"

// ElementsByTagName returns all the elements that match any of the supplied list of names.
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	addWithTag := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, tag := range name {
				if n.Data == tag {
					nodes = append(nodes, n)
					break
				}
			}
		}
	}
	forEachNode(doc, addWithTag, nil)
	return nodes
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
