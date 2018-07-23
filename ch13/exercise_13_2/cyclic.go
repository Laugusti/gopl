package cyclic

import (
	"reflect"
	"unsafe"
)

type comparison struct {
	ptr unsafe.Pointer
	t   reflect.Type
}

// IsCyclic reports whether the argument is a cyclic data structure.
func IsCyclic(v interface{}) bool {
	seen := make(map[comparison]bool)
	return isCyclic(reflect.ValueOf(&v).Elem(), seen)
}

func isCyclic(v reflect.Value, seen map[comparison]bool) bool {
	if !v.IsValid() {
		return false
	}

	if v.CanAddr() {
		ptr := unsafe.Pointer(v.UnsafeAddr())
		c := comparison{ptr, v.Type()}
		if seen[c] {
			return true // alreay seen
		}
		seen[c] = true
	}

	switch v.Kind() {
	default:
		return false

	case reflect.Ptr, reflect.Interface:
		return isCyclic(v.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if isCyclic(v.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, key := range v.MapKeys() {
			if isCyclic(v.MapIndex(key), seen) {
				return true
			}
		}
		return false
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if isCyclic(v.Field(i), seen) {
				return true
			}
		}
		return false
	}
}
