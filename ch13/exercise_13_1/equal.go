package equal

import (
	"math"
	"reflect"
	"unsafe"
)

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

// Equal reports whether x and y are deeply equal.
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	if !sameType(x, y) {
		return false
	}

	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // alreay seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return numbersCloseEnough(numberAsFloat(x), numberAsFloat(y), 1e-9)

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

		// ...struct and map cases omitted for brevity...
	case reflect.Map:
		if !equal(reflect.ValueOf(x.MapKeys()), reflect.ValueOf(y.MapKeys()), seen) {
			return false
		}
		for _, key := range x.MapKeys() {
			if !equal(x.MapIndex(key), y.MapIndex(key), seen) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

// numberAsFloat returns a float64 for v or panic if v is not a number.
func numberAsFloat(v reflect.Value) float64 {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Complex64, reflect.Complex128:
		return real(v.Complex())
	default:
		panic("not a number")
	}
}

// sameType returns true if x and y are both numbers or have the same type.
func sameType(x, y reflect.Value) bool {
	// same type
	if x.Type() == x.Type() {
		return true
	}
	// value is number if numberAsFloat does not panic
	isNumber := func(v reflect.Value) (isNum bool) {
		defer func() { isNum = recover() == nil }()
		numberAsFloat(v)
		return
	}
	// both numbers
	return isNumber(x) && isNumber(y)
}

// numbersCloseEnough returns true if the percent difference between x an y is less than pDiff, otherwise false.
func numbersCloseEnough(x, y, pDiff float64) bool {
	if x == y {
		return true
	}
	diff, avg := math.Abs(x-y), (math.Abs(x)+math.Abs(y))/2
	return diff/avg < pDiff
}
