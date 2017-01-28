package xmldom

import (
	"encoding/xml"
	"io"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

type parser struct {
	dec      *xml.Decoder
	nodeList []Node
	parents  []*Element
}

// addNode appends n as a child of the last parent or to the nodeList
func (p *parser) addNode(n Node) {
	if len(p.parents) == 0 {
		p.nodeList = append(p.nodeList, n)
	} else {
		parent := p.parents[len(p.parents)-1]
		parent.Children = append(parent.Children, n)
	}
}

// parseToken creates a Node from the token and adds it to the tree
func (p *parser) parseToken(tok xml.Token) {
	switch tok := tok.(type) {
	case xml.StartElement:
		n := &Element{Type: tok.Name, Attr: tok.Attr}
		p.addNode(n)
		// push element to parent stack
		p.parents = append(p.parents, n)
	case xml.CharData:
		p.addNode(CharData(tok))
	case xml.EndElement:
		// pop element from parent stack
		p.parents = p.parents[:len(p.parents)-1]
	}
}

// parse uses a xml.Decoder to construct a parse tree
func (p *parser) parse() error {
	for {
		tok, err := p.dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		p.parseToken(tok)
	}
	return nil
}

// Parse returns the parse tree for the XML from the given reader.
func Parse(r io.Reader) ([]Node, error) {
	p := &parser{dec: xml.NewDecoder(r)}
	if err := p.parse(); err != nil {
		return nil, err
	}
	return p.nodeList, nil
}
