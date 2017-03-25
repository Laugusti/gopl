package htmlutil

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

type document struct {
	depth  int
	source string
}

const indent = 2

// PrettyPrint returns a html.Node as a string, replacing absolute links
// in the same domain with relative links.
func PrettyPrint(url string, n *html.Node) string {
	doc := &document{}
	start := func(n *html.Node) {
		doc.startElement(getDomainFromURL(url), n)
	}
	forEachNode(n, start, doc.endElement)
	return doc.source
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

func (doc *document) startElement(domain string, n *html.Node) {
	source := doc.source
	depth := doc.depth
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
			source += getAttributes(domain, n)
		} else {
			if containsSingleTextNode(n) {
				source += fmt.Sprintf("%*s<%s", depth*indent, "", n.Data)
				source += getAttributes(domain, n)
				source += ">"
			} else {
				source += fmt.Sprintf("%*s<%s", depth*indent, "", n.Data)
				source += getAttributes(domain, n)
				source += ">\n"
			}
		}
	case n.Type == html.TextNode && !emptyOrBlankString(n.Data):
		if containsSingleTextNode(n.Parent) {
			source += strings.TrimSpace(n.Data)
		} else {
			source += getText(strings.TrimSpace(n.Data), depth)
		}
	case n.Type == html.CommentNode:
		source += fmt.Sprintf("%*s<!--", depth*indent, "")
		source += getComment(n.Data, depth)
	}
	doc.source = source
	doc.depth++
}

func (doc *document) endElement(n *html.Node) {
	doc.depth--
	source := doc.source
	depth := doc.depth
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
	doc.source = source
}

func stripDomainFromUrl(url, domain string) string {
	if strings.HasPrefix(stripSchemeFromURL(url), domain) {
		return stripSchemeFromURL(url)[len(domain):]
	}
	return url
}

func getAttributes(domain string, n *html.Node) (attributes string) {
	for _, attr := range n.Attr {
		if n.Data == "a" && attr.Key == "href" {
			url := stripDomainFromUrl(attr.Val, domain)
			attributes += fmt.Sprintf(" %s=\"%s\"", "href", url)
		} else {
			attributes += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
		}
	}
	return
}

func getText(s string, depth int) (text string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		text += fmt.Sprintf("%*s%s\n", depth*indent, "", strings.TrimSpace(line))
	}
	return
}

func getComment(s string, depth int) (comment string) {
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
