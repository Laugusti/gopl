package params

import (
	"strings"
)

type tagOptions string

// parseTag splits a struct field's http tag into it's name and
// comma-seperated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

func (o tagOptions) Options() []string {
	return strings.Split(string(o), ",")
}
