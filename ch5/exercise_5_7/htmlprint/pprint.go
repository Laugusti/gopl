package htmlprint

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

var depth int
var source string

const indent = 2

func PrettyPrint(n *html.Node) string {
	source = ""
	forEachNode(n, startElement, endElement)
	return source
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the the tree rooted at n. Both functions are optional.
// pre is called befor ethe children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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

func startElement(n *html.Node) {
	switch {
	case n.Type == html.DoctypeNode:
		source += fmt.Sprintf("<!DOCTYPE %s", n.Data)
		for _, attr := range n.Attr {
			if strings.ToLower(attr.Key) == "public" {
				source += fmt.Sprintf(" PUBLIC \"%s\"", attr.Val)
			} else {
				source += fmt.Sprintf(" \"%s\"", attr.Val)
			}
		}
		depth--
	case n.Type == html.ElementNode:
		if hasNoChildren(n) {
			source += fmt.Sprintf("%*s<%s", depth*indent, "", n.Data)
			source += getAttributes(n)
		} else {
			if containsSingleTextNode(n) {
				source += fmt.Sprintf("%*s<%s", depth*indent, "", n.Data)
				source += getAttributes(n)
				source += ">"
			} else {
				source += fmt.Sprintf("%*s<%s", depth*indent, "", n.Data)
				source += getAttributes(n)
				source += ">\n"
			}
		}
	case n.Type == html.TextNode && !emptyOrBlankString(n.Data):
		if containsSingleTextNode(n.Parent) {
			source += strings.TrimSpace(n.Data)
		} else {
			source += getText(strings.TrimSpace(n.Data))
		}
	case n.Type == html.CommentNode:
		source += fmt.Sprintf("%*s<!--", depth*indent, "")
		source += getComment(n.Data)
	}
	depth++
}

func endElement(n *html.Node) {
	depth--
	switch {
	case n.Type == html.DoctypeNode:
		source += ">\n"
	case n.Type == html.ElementNode:
		if hasNoChildren(n) {
			source += "/>\n"
		} else {
			if containsSingleTextNode(n) {
				source += fmt.Sprintf("</%s>\n", strings.TrimSpace(n.Data))
			} else {
				source += fmt.Sprintf("%*s</%s>\n", depth*indent, "", n.Data)
			}
		}
	case n.Type == html.CommentNode:
		source += "-->\n"
	}
}

func getAttributes(n *html.Node) (attributes string) {
	for _, attr := range n.Attr {
		attributes += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
	}
	return
}

func getText(s string) (text string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		text += fmt.Sprintf("%*s%s\n", depth*indent, "", strings.TrimSpace(line))
	}
	return
}

func getComment(s string) (comment string) {
	lines := strings.Split(s, "\n")
	if len(lines) > 1 {
		for _, line := range lines {
			comment += fmt.Sprintf("%*s%s\n", (depth+1)*indent, "", strings.TrimSpace(line))
		}
		comment += fmt.Sprintf("%*s", depth*indent, "")
	} else {
		comment += fmt.Sprintf(" %s  ", strings.TrimSpace(s))
	}
	return
}

func containsSingleTextNode(n *html.Node) bool {
	child := n.FirstChild
	return child != nil && child.Type == html.TextNode && strings.Index(child.Data, "\n") == -1 &&
		child.FirstChild == nil && child.NextSibling == nil
}

func hasNoChildren(n *html.Node) bool {
	return n.FirstChild == nil || (containsSingleTextNode(n) && emptyOrBlankString(n.FirstChild.Data))
}

func emptyOrBlankString(s string) bool {
	if s == "" {
		return true
	}
	for _, r := range []rune(s) {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
