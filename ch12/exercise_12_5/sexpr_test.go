package sexpr

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshalBool(t *testing.T) {
	tests := []struct {
		value bool
	}{
		{true},
		{false},
	}

	for _, test := range tests {
		b, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		var got bool
		if err := json.Unmarshal(b, &got); err != nil {
			t.Fatal(err)
		}
		if test.value != got {
			t.Errorf("TestMarshalBool => input: %v, got: %v", test.value, got)
		}
	}
}

func TestMarshalFloat(t *testing.T) {
	tests_32 := []struct {
		value float32
		want  float32
	}{
		{1, 1},
		{1.0, 1},
		{1.1, 1.100000023841858},
		{1.8233e18, 1.823300066253734e+18},
	}

	tests_64 := []struct {
		value float64
		want  float64
	}{
		{1, 1},
		{1.0, 1},
		{1.1, 1.1},
		{1.8233e18, 1.8233e+18},
	}

	for _, test := range tests_32 {
		b, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		var got float32
		if err := json.Unmarshal(b, &got); err != nil {
			t.Fatal(err)
		}
		if test.want != got {
			t.Errorf("TestMarshalFloat => input: %v, expected: %v, got: %v", test.value, test.want, got)
		}
	}
	for _, test := range tests_64 {
		b, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		var got float64
		if err := json.Unmarshal(b, &got); err != nil {
			t.Fatal(err)
		}
		if test.want != got {
			t.Errorf("TestMarshalFloat => input: %v, expected: %v, got: %v", test.value, test.want, got)
		}
	}
}

func TestMarshalComplex(t *testing.T) {
	t.Skip("complex type is not supported in json")
	tests := []struct {
		value complex64
	}{
		{3 + 2i},
		{1.1 + 2.201297e17i},
	}

	for _, test := range tests {
		b, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		var got complex64
		if err := json.Unmarshal(b, &got); err != nil {
			t.Fatal(err)
		}
		if test.value != got {
			t.Errorf("TestMarshalComplex => input: %v, got: %v", test.value, got)
		}
	}
}

func TestMarshalInterface(t *testing.T) {
	tests := []struct {
		value interface{}
	}{
		{7.2222233e89},
		{nil},
		{true},
		{"abc"},
	}

	for _, test := range tests {
		b, err := Marshal(&test.value)
		if err != nil {
			t.Fatal(err)
		}
		var got interface{}
		if err := json.Unmarshal(b, &got); err != nil {
			t.Fatal(err)
		}
		if test.value != got {
			t.Errorf("TestMarshalInterface => input: %v, got: %v", test.value, got)
		}
	}
}

func TestMarshalSlice(t *testing.T) {
	tests := []struct {
		value []int
	}{
		{[]int{1}},
		{[]int{1, 2, 3}},
	}

	for _, test := range tests {
		b, err := Marshal(&test.value)
		if err != nil {
			t.Fatal(err)
		}
		var got []int
		if err := json.Unmarshal(b, &got); err != nil {
			t.Fatal(err)
		}
		if fmt.Sprintf("%v", test.value) != fmt.Sprintf("%v", got) {
			t.Errorf("TestMarshalSlice => input: %v, got: %v", test.value, got)
		}
	}
}

func TestMarshalMap(t *testing.T) {
	t.Skip("skipping test. order not guaranteed with map")
	tests := []struct {
		value map[string]string
		want  string
	}{
		{map[string]string{"a": "1"}, `(("a" "1"))`},
		{map[string]string{"a": "1", "b": "2"}, `(("a" "1")
 ("b" "2"))`},
	}

	for _, test := range tests {
		got, err := Marshal(&test.value)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestMarshalMap => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
		}
	}
}

func TestMarshalStruct(t *testing.T) {
	t.Skip("skipping test. order not guaranteed with map")
	type roleActor struct {
		Actor map[string]string
	}
	tests := []struct {
		value roleActor
		want  string
	}{
		{roleActor{nil}, "((Actor ()))"},
		{roleActor{map[string]string{}}, "((Actor ()))"},
		{roleActor{nil}, "((Actor ()))"},
		{roleActor{map[string]string{
			"Dr. Strangelove":       "Peter Sellers",
			"Gen. Buck Turgidson":   "George C. Scott",
			`Maj. T.J. "King" Kong`: "Slim Pickens",
		}}, `((Actor (("Dr. Strangelove" "Peter Sellers")
         ("Gen. Buck Turgidson" "George C. Scott")
         ("Maj. T.J. \"King\" Kong" "Slim Pickens"))))`},
	}

	for _, test := range tests {
		got, err := Marshal(&test.value)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestMarshalStruct => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
		}
	}
}
