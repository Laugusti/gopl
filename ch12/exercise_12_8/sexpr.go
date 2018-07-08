package sexpr

import (
	"bytes"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(v); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Unmarshal decodes a S-expression into a value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	r := bytes.NewReader(data)
	if err := NewDecoder(r).Decode(v); err != nil {
		return err
	}
	return nil
}
