package intset

import (
	"testing"
)

func createIntSet(vals ...int) *IntSet {
	var x IntSet
	x.AddAll(vals...)
	return &x
}
func TestIntersectWith(t *testing.T) {
	var x IntSet
	x.IntersectWith(createIntSet())
	if x.String() != "{}" {
		t.Errorf("IntersectWith: {1 2 3} and {}, result: %v", &x)
	}

	x.AddAll(1, 2, 3)
	x.IntersectWith(createIntSet(4, 5, 6))
	if x.String() != "{}" {
		t.Errorf("IntersectWith: {1 2 3} and {4 5 6}, result: %v", &x)
	}

	x.AddAll(1, 2, 3)
	x.IntersectWith(createIntSet(1, 2, 3))
	if x.String() != "{1 2 3}" {
		t.Errorf("IntersectWith: {1 2 3} and {1 2 3 }, result: %v", &x)
	}

	x.AddAll(1, 2, 3)
	x.IntersectWith(createIntSet(1))
	if x.String() != "{1}" {
		t.Errorf("IntersectWith: {1 2 3} and {1}, result: %v", &x)
	}
}

func TestDifferenceWith(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3)
	x.DifferenceWith(createIntSet())
	if x.String() != "{1 2 3}" {
		t.Errorf("DifferenceWith: {1 2 3} and {}, result: %v", &x)
	}

	x.DifferenceWith(createIntSet(1))
	if x.String() != "{2 3}" {
		t.Errorf("DifferenceWith: {1 2 3} and {1}, result: %v", &x)
	}

	x.AddAll(1, 2, 3)
	x.DifferenceWith(createIntSet(1, 3))
	if x.String() != "{2}" {
		t.Errorf("DifferenceWith: {1 2 3} and {1 3}, result: %v", &x)
	}

	x.AddAll(1, 2, 3)
	x.DifferenceWith(createIntSet(1, 2, 3))
	if x.String() != "{}" {
		t.Errorf("DifferenceWith: {1 2 3} and {1 2 3}, result: %v", &x)
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3)
	x.SymmetricDifference(createIntSet())
	if x.String() != "{1 2 3}" {
		t.Errorf("SymmetricDifference: {1 2 3} and {}, result: %v", &x)
	}

	x.AddAll(1, 2, 3)
	x.SymmetricDifference(createIntSet(1))
	if x.String() != "{2 3}" {
		t.Errorf("SymmetricDifference: {1 2 3} and {1}, result: %v", &x)
	}

	x.AddAll(1, 2, 3)
	x.SymmetricDifference(createIntSet(1, 2))
	if x.String() != "{3}" {
		t.Errorf("SymmetricDifference: {1 2 3} and {1 2}, result: %v", &x)
	}

	x.AddAll(1, 2, 3)
	x.SymmetricDifference(createIntSet(1, 2, 3))
	if x.String() != "{}" {
		t.Errorf("SymmetricDifference: {1 2 3} and {1 2 3}, result: %v", &x)
	}

	x.AddAll(64)
	x.SymmetricDifference(createIntSet(1, 2, 3))
	if x.String() != "{1 2 3 64}" {
		t.Errorf("SymmetricDifference: {64} and {1 2 3}, result: %v", &x)
	}
}
