// ElementByID returns the Node in the document with the corresponding id attribute
package htmldom

import (
	"strings"

	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, nil)
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the the tree rooted at n. Both functions are optional.
// pre is called befor ethe children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	if pre != nil && pre(n, id) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := forEachNode(c, id, pre, post); found != nil {
			return found
		}
	}
	if post != nil && post(n, id) {
		return n
	}
	return nil
}

func startElement(n *html.Node, id string) (found bool) {
	found = getNodeID(n) == id
	return
}

func getNodeID(n *html.Node) (id string) {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if strings.ToLower(attr.Key) == "id" {
				id = attr.Val
			}
		}
	}
	return
}
