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

// An Int represents an integer in a S-expression.
type Int struct {
	Value int
}

// A Float represents a floating-point number in a S-expression.
type Float struct {
	Value float64
}

// A Complex represents a complex number in a S-expression
type Complex struct {
	Real      float64
	Imaginary float64
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
	case scanner.Float:
		f, _ := strconv.ParseFloat(dec.lex.text(), 10) // NOTE: ignoring errors
		return Float{f}, nil
	case '#':
		return dec.getComplexToken()
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

func (dec *Decoder) getComplexToken() (Token, error) {
	err := nextTokenMatches(dec, Symbol{"C"})
	if err != nil {
		return nil, err
	}
	err = nextTokenMatches(dec, StartList{})
	if err != nil {
		return nil, err
	}
	real, err := getNextTokenAsFloat64(dec)
	if err != nil {
		return nil, err
	}
	imag, err := getNextTokenAsFloat64(dec)
	if err != nil {
		return nil, err
	}
	err = nextTokenMatches(dec, EndList{})
	if err != nil {
		return nil, err
	}

	return Complex{real, imag}, nil
}

func nextTokenMatches(dec *Decoder, want Token) error {
	got, err := dec.Token()
	if err != nil {
		return err
	}
	if got != want {
		return fmt.Errorf("unexpected token %q", dec.lex.text())
	}
	return nil
}

func getNextTokenAsFloat64(dec *Decoder) (float64, error) {
	got, err := dec.Token()
	if err != nil {
		return 0, err
	}
	if got, isInt := got.(Int); isInt {
		return float64(got.Value), nil
	}
	if got, isFloat := got.(Float); isFloat {
		return got.Value, nil
	}
	return 0, fmt.Errorf("unexpected token %T, expected number", got)
}
