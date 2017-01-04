package intset

import (
	"testing"
)

func createIntSet(vals ...int) *IntSet {
	var x IntSet
	x.AddAll(vals...)
	return &x
}

func equals(x []int) bool {
	y := createIntSet(x...).Elems()
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestElems(t *testing.T) {
	if !equals([]int{}) {
		t.Fail()
	}

	if !equals([]int{1, 2, 3, 64}) {
		t.Fail()
	}
}
