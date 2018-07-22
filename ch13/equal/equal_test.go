package equal

import "testing"

func TestEqual(t *testing.T) {
	// Circular linked lists a -> b -> a and c -> c
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	tests := []struct {
		x     interface{}
		y     interface{}
		equal bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{}, []int(nil), true},
		{map[int]int{}, map[int]int(nil), true},
		{a, a, true},
		{b, b, true},
		{c, c, true},
		{a, b, false},
		{b, c, false},
	}

	for _, test := range tests {
		if Equal(test.x, test.y) != test.equal {
			t.Errorf("TestEqual inputs (%v (%[1]T) %v (%[2]T)), got %t", test.x, test.y, !test.equal)
		}
	}
}
