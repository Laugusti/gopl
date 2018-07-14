package sexpr

import (
	"reflect"
	"strings"
	"testing"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`""`, ""},
		{`"test"`, "test"},
		{`"\n"`, "\n"},
		{`"multi\nline\nstring"`, "multi\nline\nstring"},
	}

	for _, test := range tests {
		var got string
		err := NewDecoder(strings.NewReader(test.input)).Decode(&got)
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if test.want != got {
			t.Errorf("TestDecodeString => input: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}

func TestDecodeInt(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"0", 0},
		{"1", 1},
		{"2147483647", 2147483647},
		{"2147483648", 2147483648},
		{"9223372036854775807", 9223372036854775807},
	}

	for _, test := range tests {
		var got int
		err := NewDecoder(strings.NewReader(test.input)).Decode(&got)
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if test.want != got {
			t.Errorf("TestDecodeInt => input: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}

func TestDecodeIntArray(t *testing.T) {
	tests := []struct {
		input string
		want  [5]int
	}{
		{"(0 0 0 0 0)", [5]int{0, 0, 0, 0, 0}},
		{"(1 2 3 4 5)", [5]int{1, 2, 3, 4, 5}},
		{"(5 4 3 2 1)", [5]int{5, 4, 3, 2, 1}},
	}

	for _, test := range tests {
		var got [5]int
		err := NewDecoder(strings.NewReader(test.input)).Decode(&got)
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if test.want != got {
			t.Errorf("TestDecodeInt => input: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}

func TestDecodeStringSlice(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{`()`, []string{}},
		{`("a")`, []string{"a"}},
		{`("1" "2" "3" "4" "5")`, []string{"1", "2", "3", "4", "5"}},
	}

	for _, test := range tests {
		var got []string
		err := NewDecoder(strings.NewReader(test.input)).Decode(&got)
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if len(test.want) != 0 && len(got) != 0 && !reflect.DeepEqual(test.want, got) {
			t.Errorf("TestDecodeInt => input: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}

func TestDecodeStruct(t *testing.T) {
	type testStruct struct {
		A int
		B string
	}
	tests := []struct {
		input string
		want  testStruct
	}{
		{`((A 0) (B ""))`, testStruct{}},
		{`((A 1) (B "1"))`, testStruct{1, "1"}},
		{`((A 3329343) (B "this is a test"))`, testStruct{3329343, "this is a test"}},
	}

	for _, test := range tests {
		var got testStruct
		err := NewDecoder(strings.NewReader(test.input)).Decode(&got)
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)

		}
		if test.want != got {
			t.Errorf("TestDecodeStruct => input: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}

func TestDecodeMap(t *testing.T) {
	tests := []struct {
		input string
		want  map[int]string
	}{
		{`((0 ""))`, map[int]string{0: ""}},
		{`((1 "a"))`, map[int]string{1: "a"}},
		{`((2147483647 "abc\ndef\nghi\n"))`, map[int]string{2147483647: "abc\ndef\nghi\n"}},
	}

	for _, test := range tests {
		var got map[int]string
		err := NewDecoder(strings.NewReader(test.input)).Decode(&got)
		if err != nil {
			t.Fatalf("error unmarshalling %q: %v", test.input, err)
		}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("TestDecodeMap => input: %v, expected: %v, got: %v", test.input, test.want, got)
		}
	}
}
