// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type element struct {
	name  string
	attrs []xml.Attr
}

// defines regex for selector
var nameRegex = regexp.MustCompile(`^\w+`)
var idRegex = regexp.MustCompile(`#\w+`)
var classRegex = regexp.MustCompile(`\.\w+`)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []element // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, element{tok.Name.Local, tok.Attr}) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", getNames(stack), tok)
			}
		}
	}
}

// getNames returns all the names in the element slice
func getNames(elements []element) string {
	var names string
	for i, el := range elements {
		if i != 0 {
			names += " "
		}
		names += el.name
	}
	return names
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []element, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if matches(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// matches returns true if element matches the selector
func matches(el element, s string) bool {
	selector := getSelectorMap(s)
	if name, ok := selector["elementName"]; ok {
		if el.name != name {
			return false
		}
	}
	if id, ok := selector["id"]; ok {
		if getAttributeValue(el.attrs, "id") != id {
			return false
		}
	}
	if class, ok := selector["class"]; ok {
		hasClass := false
		for _, c := range strings.Split(getAttributeValue(el.attrs, "class"), " ") {
			if c == class {
				hasClass = true
			}
		}
		if !hasClass {
			return false
		}
	}
	return true
}

// getSelectorMap creates map of element attributes to their value from a string
func getSelectorMap(s string) map[string]string {
	selector := map[string]string{}
	if name := nameRegex.FindString(s); name != "" {
		selector["elementName"] = name
	}
	if id := idRegex.FindString(s); id != "" {
		selector["id"] = id[1:]
	}
	if class := classRegex.FindString(s); class != "" {
		selector["class"] = class[1:]
	}
	return selector
}

// getAttributeValue returns the Value of the specified attribute
func getAttributeValue(attrs []xml.Attr, a string) string {
	for _, attr := range attrs {
		if attr.Name.Local == a {
			return attr.Value
		}
	}
	return ""
}
