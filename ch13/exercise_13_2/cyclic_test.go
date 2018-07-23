package cyclic

import "testing"
import "unsafe"

func TestIsCyclic(t *testing.T) {
	// Circular linked lists a -> b -> a and c -> c
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	tests := []struct {
		input interface{}
		cycle bool
	}{
		{0, false},
		{[]int{1, 2, 3}, false},
		{a, true},
		{b, true},
		{c, true},
		{*c, true},
		{&c, true},
		{[]interface{}{a}, true},
		{[3]interface{}{a, b, c}, true},
		{map[int]link{2: *b}, true},
		{unsafe.Pointer(&a), false},
	}

	for _, test := range tests {
		if IsCyclic(test.input) != test.cycle {
			t.Errorf("TestIsCyclic inputs (%v), got %t", test.input, !test.cycle)
		}
	}
}
