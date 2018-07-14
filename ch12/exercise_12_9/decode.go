package sexpr

import (
	"fmt"
	"io"
	"reflect"
	"text/scanner"
)

// A Decoder reads an decodes a S-Expression from an input stream.
type Decoder struct {
	r   io.Reader
	lex *lexer
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

// Decode reads the next S-Expression value from it's input
// and stores it in the value pointed to by v.
func (dec *Decoder) Decode(v interface{}) (err error) {
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.lex.scan.Position, x)
		}
	}()
	setValueFromTokens(reflect.ValueOf(v).Elem(), dec.getAllTokens())
	return nil
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func (dec *Decoder) getAllTokens() []Token {
	var tokens []Token
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tokens = append(tokens, tok)
	}
	return tokens
}

func setValueFromTokens(v reflect.Value, tokens []Token) {
	switch {
	case len(tokens) == 1:
		setValue(v, tokens[0])
	case len(tokens) > 1:
		setListValue(v, tokens)
	default:
		panic("no tokens read")
	}
}

func setValue(v reflect.Value, token Token) {
	switch token := token.(type) {
	case Symbol:
		if token.Name == "nil" {
			v.Set(reflect.Zero(v.Type()))
		}
	case Int:
		v.SetInt(int64(token.Value))
	case String:
		v.SetString(token.Text)
	default:
		panic(fmt.Sprintf("unexpected token %T", token))
	}
}

func setListValue(v reflect.Value, tokens []Token) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i, indexTokens := range getListAsTokenList(tokens) {
			setValueFromTokens(v.Index(i), indexTokens)
		}
	case reflect.Slice: // (item ...)
		for _, indexTokens := range getListAsTokenList(tokens) {
			item := reflect.New(v.Type().Elem()).Elem()
			setValueFromTokens(item, indexTokens)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct: // ((name value) ...)
		for _, indexTokens := range getListAsTokenList(tokens) {
			nameValuePair := getListAsTokenList(indexTokens)
			if _, isSymbol := nameValuePair[0][0].(Symbol); !isSymbol {
				panic(fmt.Sprintf("got token %T, want field Symbol", nameValuePair[0]))
			}
			name := nameValuePair[0][0].(Symbol).Name
			setValueFromTokens(v.FieldByName(name), nameValuePair[1])
		}
	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for _, indexTokens := range getListAsTokenList(tokens) {
			nameValuePair := getListAsTokenList(indexTokens)
			key := reflect.New(v.Type().Key()).Elem()
			setValueFromTokens(key, nameValuePair[0])
			value := reflect.New(v.Type().Elem()).Elem()
			setValueFromTokens(value, nameValuePair[1])
			v.SetMapIndex(key, value)
		}
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func getListAsTokenList(tokens []Token) [][]Token {
	if _, start := tokens[0].(StartList); !start {
		panic(fmt.Sprintf("unexpected token %T", tokens[0]))
	}
	if _, end := tokens[len(tokens)-1].(EndList); !end {
		panic("end of file")
	}
	tokens = tokens[1 : len(tokens)-1]
	var tokensList [][]Token
	count := 0
	curElementStartIndex := 0
	for i, tok := range tokens {
		switch tok.(type) {
		case StartList:
			count++
		case EndList:
			count--
		}
		if count <= 0 {
			tokensList = append(tokensList, tokens[curElementStartIndex:i+1])
			curElementStartIndex = i + 1
		}
	}
	return tokensList
}
