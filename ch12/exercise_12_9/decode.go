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
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		setValueFromToken(dec, tok, reflect.ValueOf(v).Elem())
	}
	return nil
}

func setValueFromToken(dec *Decoder, tok Token, v reflect.Value) {
	switch tok := tok.(type) {
	case Symbol:
		if tok.Name == "nil" {
			v.Set(reflect.Zero(v.Type()))
		}
	case Int:
		v.SetInt(int64(tok.Value))
	case String:
		v.SetString(tok.Text)
	case StartList:
		decodeList(dec, v)
	default:
		panic(fmt.Sprintf("unexpected token %T", tok))
	}
}

func decodeList(dec *Decoder, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; ; i++ {
			tok, endOfList := getToken(dec)
			if endOfList {
				break
			}
			setValueFromToken(dec, tok, v.Index(i))
		}
	case reflect.Slice: // (item ...)
		for i := 0; ; i++ {
			tok, endOfList := getToken(dec)
			if endOfList {
				break
			}
			item := reflect.New(v.Type().Elem()).Elem()
			setValueFromToken(dec, tok, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct: // ((name value) ...)
		for i := 0; ; i++ {
			tok, endOfList := getToken(dec)
			if endOfList {
				break
			}
			if _, isStartList := tok.(StartList); !isStartList {
				panic(fmt.Sprintf("wanted StartList, got %T", tok))
			}
			tok, _ = getToken(dec)
			if _, isSymbol := tok.(Symbol); !isSymbol {
				panic(fmt.Sprintf("wanted Symbol, got %T", tok))
			}
			name := tok.(Symbol).Name
			tok, _ = getToken(dec)
			setValueFromToken(dec, tok, v.FieldByName(name))
			tok, _ = getToken(dec)
			if _, isEnd := tok.(EndList); !isEnd {
				panic(fmt.Sprintf("wanted EndList, got %T", tok))
			}
		}
	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for i := 0; ; i++ {
			tok, endOfList := getToken(dec)
			if endOfList {
				break
			}
			if _, isStartList := tok.(StartList); !isStartList {
				panic(fmt.Sprintf("wanted StartList, got %T", tok))
			}
			key := reflect.New(v.Type().Key()).Elem()
			tok, _ = getToken(dec)
			setValueFromToken(dec, tok, key)
			value := reflect.New(v.Type().Elem()).Elem()
			tok, _ = getToken(dec)
			setValueFromToken(dec, tok, value)
			v.SetMapIndex(key, value)
			tok, _ = getToken(dec)
			if _, isEnd := tok.(EndList); !isEnd {
				panic(fmt.Sprintf("wanted EndList, got %T", tok))
			}
		}
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}
func getToken(dec *Decoder) (Token, bool) {
	tok, err := dec.Token()
	if err == io.EOF {
		panic("end of file")
	}
	if err != nil {
		panic(err)
	}
	_, isEnd := tok.(EndList)
	return tok, isEnd
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
