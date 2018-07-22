package sexpr

import "reflect"
import "testing"

func TestEncodeNameTag(t *testing.T) {
	// - is not a symbol
	data := struct {
		A int `sexpr:"a"`
		B int `sexpr:"-,"`
	}{1, 2}
	got, err := Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	want := "((a 1)\n (- 2))"
	if want != string(got) {
		t.Errorf("TestEncodeNameTag => input: %v, expected: %v, got: %v", data, want, string(got))
	}
}

func TestDecodeNameTag(t *testing.T) {
	// - is not a symbol
	type testStruct struct {
		A int `sexpr:"a"`
		B int `sexpr:"z,"`
	}
	b := []byte("((a 1)\n (z 2))")
	var got testStruct
	err := Unmarshal(b, &got)
	if err != nil {
		t.Fatal(err)
	}
	want := testStruct{1, 2}
	if got != want {
		t.Errorf("TestDecodeNameTag => input: %v, expected: %v, got: %v", string(b), want, got)
	}
}

func TestEncodeOmitEmptyTag(t *testing.T) {
	type testStruct struct {
		A int    `sexpr:",omitempty"`
		B string `sexpr:",omitempty"`
	}
	tests := []struct {
		input testStruct
		want  string
	}{
		{testStruct{0, ""}, "()"},
		{testStruct{1, ""}, "((A 1))"},
		{testStruct{0, "a"}, "((B \"a\"))"},
		{testStruct{1, "a"}, "((A 1)\n (B \"a\"))"},
	}

	for _, test := range tests {
		got, err := Marshal(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestEncodeOmitEmptyTag => input: %v, expected: %v, got: %v", test.input, test.want, string(got))
		}
	}
}

func TestEncodeAlwaysOmitTag(t *testing.T) {
	type testStruct struct {
		A int
		B string `sexpr:"-"`
	}
	tests := []struct {
		input testStruct
		want  string
	}{
		{testStruct{0, ""}, "((A 0))"},
		{testStruct{1, ""}, "((A 1))"},
		{testStruct{0, "a"}, "((A 0))"},
		{testStruct{1, "a"}, "((A 1))"},
	}

	for _, test := range tests {
		got, err := Marshal(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestEncodeAlwaysOmitTag => input: %v, expected: %v, got: %v", test.input, test.want, string(got))
		}
	}
}

func TestEncodeStringTag(t *testing.T) {
	type testStruct struct {
		A bool        `sexpr:",string"`
		B int         `sexpr:",string"`
		C float64     `sexpr:",string"`
		D string      `sexpr:",string"`
		E map[int]int `sexpr:",string"`
	}
	tests := []struct {
		input testStruct
		want  string
	}{
		{testStruct{}, `((A "nil")
 (B "0")
 (C "0.00000")
 (D "\"\"")
 (E ()))`},
	}

	for _, test := range tests {
		got, err := Marshal(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestEncodeStringTag => input: %v, expected: %v, got: %v", test.input, test.want, string(got))
		}
	}
}

func TestDecodeStringTag(t *testing.T) {
	type testStruct struct {
		A bool    `sexpr:",string"`
		B int     `sexpr:",string"`
		C float64 `sexpr:",string"`
		D string  `sexpr:",string"`
	}
	tests := []struct {
		input string
		want  testStruct
	}{
		{`((A "nil")
 (B "0")
 (C "0.0")
 (D "\"\""))`, testStruct{}},
	}

	for _, test := range tests {
		var got testStruct
		err := Unmarshal([]byte(test.input), &got)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("TestDecodeStringTag => input: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}
