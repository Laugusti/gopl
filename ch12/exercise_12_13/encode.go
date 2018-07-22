package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// An Encoder writes a S-Expressions to an output stream
type Encoder struct {
	w io.Writer
}

// NewEncoder returns a new encoder that writes to w
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w}
}

// Encode writes the S-Expression encoding of v to the stream
func (enc *Encoder) Encode(v interface{}) error {
	return enc.encode(reflect.ValueOf(v), "")
}

func (enc *Encoder) encode(v reflect.Value, prefix string) error {
	switch v.Kind() {
	case reflect.Invalid:
		if _, err := fmt.Fprintf(enc.w, "nil"); err != nil {
			return err
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if _, err := fmt.Fprintf(enc.w, "%d", v.Int()); err != nil {
			return err
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if _, err := fmt.Fprintf(enc.w, "%d", v.Uint()); err != nil {
			return err
		}

	case reflect.String:
		if _, err := fmt.Fprintf(enc.w, "%q", v.String()); err != nil {
			return err
		}

	case reflect.Ptr:
		return enc.encode(v.Elem(), prefix)

	case reflect.Array, reflect.Slice: // (value ...)
		if _, err := fmt.Fprintf(enc.w, "("); err != nil {
			return err
		}
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				if _, err := fmt.Fprintf(enc.w, "\n%s ", prefix); err != nil {
					return err
				}
			}
			if err := enc.encode(v.Index(i), prefix+" "); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(enc.w, ")"); err != nil {
			return err
		}

	case reflect.Struct: // ((name value) ...)
		if _, err := fmt.Fprintf(enc.w, "("); err != nil {
			return err
		}
		addLine := false
		for i := 0; i < v.NumField(); i++ {
			tag := getFieldTag("sexpr", v.Type().Field(i))
			// skip if tag name is "-" and no options or tag contains
			// "omitempty" and field is zero value
			if (tag.Name == "-" && !tag.HasOptions) || (isZeroValue(v.Field(i)) && tag.Contains("omitempty")) {
				continue
			}
			if addLine {
				if _, err := fmt.Fprintf(enc.w, "\n%s ", prefix); err != nil {
					return err
				}
			} else {
				addLine = true
			}
			if _, err := fmt.Fprintf(enc.w, "(%s ", tag.Name); err != nil {
				return err
			}
			// string option encodes ints, floats and bools as string
			if tag.Contains("string") && isPrimitive(v.Field(i)) {
				s, err := encodeAsQuotedString(v.Field(i))
				if err != nil {
					return err
				}
				if _, err := fmt.Fprint(enc.w, s); err != nil {
					return err
				}
			} else {
				newPrefix := prefix + strings.Repeat(" ", len(tag.Name)+3)
				if err := enc.encode(v.Field(i), newPrefix); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprintf(enc.w, ")"); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(enc.w, ")"); err != nil {
			return err
		}

	case reflect.Map: // ((key value) ...)
		if _, err := fmt.Fprintf(enc.w, "("); err != nil {
			return err
		}
		for i, key := range v.MapKeys() {
			if i > 0 {
				if _, err := fmt.Fprintf(enc.w, "\n%s ", prefix); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprintf(enc.w, "("); err != nil {
				return err
			}
			tmpEnc := NewEncoder(&bytes.Buffer{})
			if err := tmpEnc.encode(key, prefix+" "); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(enc.w, tmpEnc.w.(*bytes.Buffer).String()); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(enc.w, " "); err != nil {
				return err
			}
			newPrefix := prefix + strings.Repeat(" ", len(tmpEnc.w.(*bytes.Buffer).String())+2)
			if err := enc.encode(v.MapIndex(key), newPrefix); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(enc.w, ")"); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(enc.w, ")"); err != nil {
			return err
		}
	case reflect.Bool:
		if v.Bool() {
			if _, err := fmt.Fprintf(enc.w, "t"); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(enc.w, "nil"); err != nil {
				return err
			}
		}

	case reflect.Float32, reflect.Float64:
		if _, err := fmt.Fprintf(enc.w, "%#g", v.Float()); err != nil {
			return err
		}

	case reflect.Complex64, reflect.Complex128:
		if _, err := fmt.Fprintf(enc.w, "#C(%#g %#g)", real(v.Complex()), imag(v.Complex())); err != nil {
			return err
		}

	case reflect.Interface:
		if v.IsNil() {
			if _, err := fmt.Fprintf(enc.w, "nil"); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(enc.w, "("); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(enc.w, "%q ", v.Elem().Type()); err != nil {
				return err
			}
			newPrefix := prefix + strings.Repeat(" ", len(v.Elem().Type().String())+2)
			if err := enc.encode(v.Elem(), newPrefix); err != nil {
				return err
			}
			if _, err := fmt.Fprintf(enc.w, ")"); err != nil {
				return err
			}
		}

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

// encodeAsQuotedString returns the encoded value as a quoted string
func encodeAsQuotedString(v reflect.Value) (string, error) {
	b, err := Marshal(v.Interface())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%q", b), nil

}

// isPrimitive returns true for bools, ints, floats, and strings, false otherwise
func isPrimitive(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fallthrough
	case reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.String:
		return true
	default:
		return false
	}

}
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fallthrough

	case reflect.String:
		fallthrough

	case reflect.Bool:
		fallthrough

	case reflect.Float32, reflect.Float64:
		fallthrough

	case reflect.Complex64, reflect.Complex128:
		return v.Interface() == reflect.Zero(v.Type()).Interface()

	case reflect.Interface:
		return v.IsNil()

	case reflect.Ptr:
		return v.Elem().Kind() == reflect.Invalid

	case reflect.Array, reflect.Slice:
		return v.Len() == 0

	case reflect.Map:
		return len(v.MapKeys()) == 0

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZeroValue(v.Field(i)) {
				return false
			}
		}
		return true

	default:
		return false
	}
}
