package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	var x IntSet
	if x.Len() != 0 {
		t.Fail()
	}

	x.Add(1)
	if x.Len() != 1 {
		t.Fail()
	}

	x.Add(1)
	if x.Len() != 1 {
		t.Fail()
	}

	x.Add(2)
	if x.Len() != 2 {
		t.Fail()
	}

	x.Add(11000)
	if x.Len() != 3 {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Remove(1)
	if x.Has(1) || x.Len() != 0 {
		t.Fail()
	}

	x.Add(1)
	x.Add(2)
	x.Add(11000)

	x.Remove(123)
	if x.Len() != 3 {
		t.Fail()
	}

	x.Remove(2)
	if x.Len() != 2 {
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(11000)

	x.Clear()
	if x.Len() != 0 {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(11000)

	copy := *x.Copy()
	if x.String() != copy.String() {
		t.Fail()
	}
}
