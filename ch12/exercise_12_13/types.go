package sexpr

import "fmt"
import "reflect"
import "strconv"
import "strings"

type typeKind int

const (
	basic = iota
	ptr
	array
	slice
)

var (
	builtinTypes = map[string]reflect.Type{
		// boolean
		"bool": reflect.TypeOf(true),
		// byte, rune, string
		"byte":   reflect.TypeOf(byte(1)),
		"rune":   reflect.TypeOf(rune(1)),
		"string": reflect.TypeOf(""),
		// integers
		"int":     reflect.TypeOf(int(1)),
		"int8":    reflect.TypeOf(int8(1)),
		"int16":   reflect.TypeOf(int16(1)),
		"int32":   reflect.TypeOf(int32(1)),
		"int64":   reflect.TypeOf(int32(1)),
		"uint":    reflect.TypeOf(uint(1)),
		"uint8":   reflect.TypeOf(uint8(1)),
		"uint16":  reflect.TypeOf(uint16(1)),
		"uint32":  reflect.TypeOf(uint32(1)),
		"uint64":  reflect.TypeOf(uint64(1)),
		"uintptr": reflect.TypeOf(uintptr(1)),
		// floats
		"float32": reflect.TypeOf(float32(1)),
		"float64": reflect.TypeOf(float64(1)),
		// complex
		"complex64":  reflect.TypeOf(complex64(complex(0, 1))),
		"complex128": reflect.TypeOf(complex128(complex(0, 1))),
	}
	customTypes = map[string]reflect.Type{}
)

// RegisterCustomType registers name as a custom type (using an example value) when decoding an interface value.
func RegisterCustomType(name string, v interface{}) {
	customTypes[name] = reflect.TypeOf(v)
}

func getTypeKind(name string) typeKind {
	switch {
	case strings.HasPrefix(name, "*"):
		return ptr
	case strings.HasPrefix(name, "[]"):
		return slice
	case strings.HasPrefix(name, "["):
		return array
	default:
		return basic
	}
}

func getTypeForName(name string) reflect.Type {
	switch getTypeKind(name) {
	case ptr:
		return reflect.PtrTo(getTypeForName(strings.TrimPrefix(name, "*")))
	case slice:
		return reflect.SliceOf(getTypeForName(strings.TrimPrefix(name, "[]")))
	case array:
		size, err := strconv.Atoi(name[1:strings.IndexRune(name, ']')])
		if err != nil {
			break
		}
		return reflect.ArrayOf(size, getTypeForName(name[strings.IndexRune(name, ']')+1:]))
	case basic:
		if t, ok := builtinTypes[name]; ok {
			return t
		}
		if t, ok := customTypes[name]; ok {
			return t
		}
	}
	panic(fmt.Sprintf("cannot decode to type %q", name))
}
