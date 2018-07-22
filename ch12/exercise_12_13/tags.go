package sexpr

import (
	"reflect"
	"strings"
)

type fieldTag struct {
	Name       string
	IsTagged   bool
	HasOptions bool
	Options    []string
}

func getFieldTag(tag string, fieldInfo reflect.StructField) fieldTag {
	t := fieldInfo.Tag.Get(tag)
	if idx := strings.Index(t, ","); idx != -1 {
		name := t[:idx]
		if name == "" {
			name = fieldInfo.Name
		}
		return fieldTag{name, true, true, strings.Split(t[idx+1:], ",")}
	}
	if t == "" {
		return fieldTag{fieldInfo.Name, false, false, nil}
	}
	return fieldTag{t, true, false, nil}
}

func (t fieldTag) Contains(s string) bool {
	if !t.HasOptions {
		return false
	}
	for _, opt := range t.Options {
		if s == opt {
			return true
		}
	}
	return false
}
