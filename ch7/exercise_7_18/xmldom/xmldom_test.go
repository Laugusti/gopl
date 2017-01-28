package xmldom

import (
	"fmt"
	"strings"
	"testing"
)

func nodeListAsString(nodes []Node) string {
	var s string
	for _, n := range nodes {
		switch n := n.(type) {
		case CharData:
			s += string(n)
		case *Element:
			s += fmt.Sprintf("<%s", n.Type.Local)
			for _, a := range n.Attr {
				s += fmt.Sprintf(" %s=\"%s\"", a.Name.Local, a.Value)
			}
			s += ">"
			s += nodeListAsString(n.Children)
			s += fmt.Sprintf("</%s>", n.Type.Local)
		}
	}
	return s
}
func TestParse(t *testing.T) {
	tests := []struct {
		data string
		want string
	}{
		{"<a><b>Test</b><c/></a>", "<a><b>Test</b><c></c></a>"},
		{`<a href="http://abc.com">abc</a>`, `<a href="http://abc.com">abc</a>`},
		{"data", "data"},
		{"data<a/>", "data<a></a>"},
	}

	for _, test := range tests {
		data := strings.NewReader(test.data)
		nodes, err := Parse(data)
		if err != nil {
			t.Errorf("parse error: %v", err)
		}
		if nodeListAsString(nodes) != test.want {
			t.Errorf("nodeListAsString = %q, want %q",
				nodeListAsString(nodes), test.want)
		}
	}
}
