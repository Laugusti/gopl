package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func encode(buf *bytes.Buffer, v reflect.Value, prefix string) error {
	if isZeroValue(v) {
		return nil
	}
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), prefix)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%s ", prefix)
			}
			if err := encode(buf, v.Index(i), prefix+" "); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				fmt.Fprintf(buf, "\n%s ", prefix)
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			newPrefix := prefix + strings.Repeat(" ", len(v.Type().Field(i).Name)+3)
			if err := encode(buf, v.Field(i), newPrefix); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				fmt.Fprintf(buf, "\n%s ", prefix)
			}
			buf.WriteByte('(')
			var tmpBuf bytes.Buffer
			if err := encode(&tmpBuf, key, prefix+" "); err != nil {
				return err
			}
			buf.Write(tmpBuf.Bytes())
			buf.WriteByte(' ')
			newPrefix := prefix + strings.Repeat(" ", len(tmpBuf.String())+2)
			if err := encode(buf, v.MapIndex(key), newPrefix); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%v", v.Float())

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%v %v)", real(v.Complex()), imag(v.Complex()))

	case reflect.Interface:
		if v.IsNil() {
			buf.WriteString("nil")
		} else {
			buf.WriteByte('(')
			fmt.Fprintf(buf, "%q ", v.Elem().Type())
			newPrefix := prefix + strings.Repeat(" ", len(v.Elem().Type().String())+2)
			if err := encode(buf, v.Elem(), newPrefix); err != nil {
				return err
			}
			buf.WriteByte(')')
		}

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
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

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), ""); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
