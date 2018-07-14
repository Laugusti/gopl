package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

// A Token is an inteface holding one  of the token types:
// StartList, EndList, Symbol, String, Int.
type Token interface{}

// A Symbol represents valid identifiers (nil and struct field names) in a S-expression.
type Symbol struct {
	Name string
}

// A String respresents raw text in a S-expresion.
type String struct {
	Text string
}

// An Int represents represents an integer in a S-expression.
type Int struct {
	Value int
}

// StartList represents the start of a S-expression list.
type StartList struct {
}

// EndList represents the end of a S-expression list.
type EndList struct {
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

// Token returns the next S-expression token in the input stream.
// At the end of the input stream, Token returns nil, io.EOF.
func (dec *Decoder) Token() (Token, error) {
	if dec.lex == nil {
		dec.lex = &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
		dec.lex.scan.Init(dec.r)
	}
	// get token
	dec.lex.next()
	switch dec.lex.token {
	case scanner.Ident:
		return Symbol{dec.lex.text()}, nil
	case scanner.String:
		s, _ := strconv.Unquote(dec.lex.text()) // NOTE: ignoring errors
		return String{s}, nil
	case scanner.Int:
		i, _ := strconv.Atoi(dec.lex.text()) // NOTE: ignoring errors
		return Int{i}, nil
	case '(':
		return StartList{}, nil
	case ')':
		return EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected token %q", dec.lex.text())
	}
}
