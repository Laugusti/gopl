package sexpr

import (
	"strings"
	"testing"
)

func TestStartListToken(t *testing.T) {
	tests := []struct {
		input string
		want  Token
	}{
		{"(", StartList{}},
	}
	for _, test := range tests {
		got, err := NewDecoder(strings.NewReader(test.input)).Token()
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("TestEndListToken => input: %q, expected: %[2]T%[2]v, got %[3]T%[3]v", test.input, test.want, got)
		}
	}
}

func TestEndListToken(t *testing.T) {
	tests := []struct {
		input string
		want  Token
	}{
		{")", EndList{}},
	}
	for _, test := range tests {
		got, err := NewDecoder(strings.NewReader(test.input)).Token()
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("TestEndListToken => input: %q, expected: %[2]T%[2]v, got %[3]T%[3]v", test.input, test.want, got)
		}
	}
}

func TestSymbolToken(t *testing.T) {
	tests := []struct {
		input string
		want  Token
	}{
		{"nil", Symbol{"nil"}},
		{"t", Symbol{"t"}},
		{"value", Symbol{"value"}},
	}
	for _, test := range tests {
		got, err := NewDecoder(strings.NewReader(test.input)).Token()
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("TestSymbolToken => input: %q, expected: %[2]T%[2]v, got %[3]T%[3]v", test.input, test.want, got)
		}
	}
}

func TestStringToken(t *testing.T) {
	tests := []struct {
		input string
		want  Token
	}{
		{`""`, String{""}},
		{`"nil"`, String{"nil"}},
		{`"this is a string"`, String{"this is a string"}},
	}
	for _, test := range tests {
		got, err := NewDecoder(strings.NewReader(test.input)).Token()
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("TestStringToken => input: %q, expected: %[2]T%[2]v, got %[3]T%[3]v", test.input, test.want, got)
		}
	}
}

func TestIntToken(t *testing.T) {
	tests := []struct {
		input string
		want  Token
	}{
		{"0", Int{0}},
		{"1", Int{1}},
		{"2147483647", Int{2147483647}},
		{"9223372036854775807", Int{9223372036854775807}},
	}
	for _, test := range tests {
		got, err := NewDecoder(strings.NewReader(test.input)).Token()
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("TestIntToken => input: %q, expected: %[2]T%[2]v, got %[3]T%[3]v", test.input, test.want, got)
		}
	}
}

func TestFloatToken(t *testing.T) {
	tests := []struct {
		input string
		want  Token
	}{
		{"0.0", Float{0}},
		{"3.40282346638528859811704183484516925440e+38", Float{3.40282346638528859811704183484516925440e+38}},
		{"1.797693134862315708145274237317043567981e+308", Float{1.797693134862315708145274237317043567981e+308}},
	}
	for _, test := range tests {
		got, err := NewDecoder(strings.NewReader(test.input)).Token()
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("TestFloatToken => input: %q, expected: %[2]T%[2]v, got %[3]T%[3]v", test.input, test.want, got)
		}
	}
}

func TestComplexToken(t *testing.T) {
	tests := []struct {
		input string
		want  Token
	}{
		{"#C(0 0)", Complex{0, 0}},
		{"#C(1 1)", Complex{1, 1}},
		{"#C(1.1 2.201297e17)", Complex{1.1, 2.201297e+17}},
	}
	for _, test := range tests {
		got, err := NewDecoder(strings.NewReader(test.input)).Token()
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("TestComplexToken => input: %q, expected: %[2]T%[2]v, got %[3]T%[3]v", test.input, test.want, got)
		}
	}
}
