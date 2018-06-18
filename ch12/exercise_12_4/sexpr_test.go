package sexpr

import (
	"testing"
)

func TestMarshalBool(t *testing.T) {
	tests := []struct {
		value bool
		want  string
	}{
		{true, "t"},
		{false, "nil"},
	}

	for _, test := range tests {
		got, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestMarshalBool => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
		}
	}
}

func TestMarshalFloat(t *testing.T) {
	tests_32 := []struct {
		value float32
		want  string
	}{
		{1, "1"},
		{1.0, "1"},
		{1.1, "1.100000023841858"},
		{1.8233e18, "1.823300066253734e+18"},
	}

	tests_64 := []struct {
		value float64
		want  string
	}{
		{1, "1"},
		{1.0, "1"},
		{1.1, "1.1"},
		{1.8233e18, "1.8233e+18"},
	}

	for _, test := range tests_32 {
		got, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestMarshalBool => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
		}
	}
	for _, test := range tests_64 {
		got, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestMarshalBool => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
		}
	}
}

func TestMarshalComplex(t *testing.T) {
	tests := []struct {
		value complex64
		want  string
	}{
		{3 + 2i, "#C(3 2)"},
		{1.1 + 2.201297e17i, "#C(1.100000023841858 2.2012970112385024e+17)"},
	}

	for _, test := range tests {
		got, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestMarshalBool => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
		}
	}
}

func TestMarshalInterface(t *testing.T) {
	tests := []struct {
		value interface{}
		want  string
	}{
		{1, `("int" 1)`},
		{7.2222233e89, `("float64" 7.2222233e+89)`},
		{nil, `nil`},
		{true, `("bool" t)`},
		{"abc", `("string" "abc")`},
		{[]int{1, 2, 3}, "(\"[]int\" (1\n        2\n        3))"},
	}

	for _, test := range tests {
		got, err := Marshal(&test.value)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestMarshalBool => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
		}
	}
}

func TestMarshalSlice(t *testing.T) {
	tests := []struct {
		value []int
		want  string
	}{
		{[]int{1}, "(1)"},
		{[]int{1, 2, 3}, "(1\n 2\n 3)"},
	}

	for _, test := range tests {
		got, err := Marshal(&test.value)
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Errorf("TestMarshalBool => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
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
			t.Errorf("TestMarshalBool => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
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
			t.Errorf("TestMarshalBool => input: %v, expected: %v, got: %v", test.value, test.want, string(got))
		}
	}
}
