package sexpr

import (
	"fmt"
	"io"
	"reflect"
)

// A Decoder reads a S-Expression from an input stream.
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
		} else if tok.Name == "t" {
			v.SetBool(true)
		}
	case Int:
		v.SetInt(int64(tok.Value))
	case Float:
		v.SetFloat(tok.Value)
	case Complex:
		v.SetComplex(complex(tok.Real, tok.Imaginary))
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
		doUntilEndofList(dec, func(i int, tok Token) {
			setValueFromToken(dec, tok, v.Index(i))
		})
	case reflect.Slice: // (item ...)
		doUntilEndofList(dec, func(i int, tok Token) {
			item := reflect.New(v.Type().Elem()).Elem()
			setValueFromToken(dec, tok, item)
			v.Set(reflect.Append(v, item))
		})
	case reflect.Struct: // ((name value) ...)
		doUntilEndofList(dec, func(i int, tok Token) {
			typeMustMatch(StartList{}, tok)

			tok, _ = getToken(dec)
			typeMustMatch(Symbol{}, tok)
			name := tok.(Symbol).Name
			tok, _ = getToken(dec)
			if field := getFieldByTagName(v, name); field.IsValid() {
				// field has "string" option, we need to decode first
				if structFieldTag := getFieldTag("sexpr", v.Type().Field(i)); structFieldTag.Contains("string") {
					if !isPrimitive(field) {
						panic(fmt.Sprintf("cannot use field of type %q with string option", field.Type().Name()))
					}
					err := Unmarshal([]byte(tok.(String).Text), field.Addr().Interface())
					if err != nil {
						panic(err)
					}
				} else {
					setValueFromToken(dec, tok, field)
				}
			} else {
				panic(fmt.Sprintf("cannot find single field with name/tag %q in %s", name, v.Type().Name()))
			}

			tok, _ = getToken(dec)
			typeMustMatch(EndList{}, tok)
		})
	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		doUntilEndofList(dec, func(i int, tok Token) {
			typeMustMatch(StartList{}, tok)

			key := reflect.New(v.Type().Key()).Elem()
			tok, _ = getToken(dec)
			setValueFromToken(dec, tok, key)
			value := reflect.New(v.Type().Elem()).Elem()
			tok, _ = getToken(dec)
			setValueFromToken(dec, tok, value)
			v.SetMapIndex(key, value)

			tok, _ = getToken(dec)
			typeMustMatch(EndList{}, tok)
		})
	case reflect.Interface: // (type value)
		// get type and create new value of type
		tok, _ := getToken(dec)
		typeMustMatch(String{}, tok)
		typeName := tok.(String).Text
		value := reflect.New(getTypeForName(typeName)).Elem()

		// read value into created value
		tok, _ = getToken(dec)
		setValueFromToken(dec, tok, value)

		// set interface using value
		v.Set(value)

		tok, _ = getToken(dec)
		typeMustMatch(EndList{}, tok)
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

func doUntilEndofList(dec *Decoder, f func(int, Token)) {
	for i := 0; ; i++ {
		tok, endOfList := getToken(dec)
		if endOfList {
			break
		}
		f(i, tok)
	}
}

func typeMustMatch(want, got interface{}) {
	if reflect.TypeOf(want) != reflect.TypeOf(got) {
		panic(fmt.Sprintf("wanted %T, got %T", want, got))
	}
}

// getFieldByTagName returns the field value in the struct with the tagName
// or an Invalid value if none found
func getFieldByTagName(v reflect.Value, tagName string) reflect.Value {
	// default reflect.Value is not valid
	var field reflect.Value
	var gotTaggedField bool
	for i := 0; i < v.NumField(); i++ {
		tag := getFieldTag("sexpr", v.Type().Field(i))
		if tag.Name == tagName {
			// return default reflect.Value if multiple fields with
			// same tag
			if tag.IsTagged && gotTaggedField {
				return reflect.Value{}
			}
			// prefer tagged fields
			if !gotTaggedField {
				field = v.Field(i)
				gotTaggedField = tag.IsTagged
			}
		}
	}
	return field
}
