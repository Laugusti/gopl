package geometry

import (
	"testing"
)

func TestPointDistance(t *testing.T) {
	p, q := Point{1, 2}, Point{4, 6}

	if Distance(p, q) != 5 {
		t.Fail()
	}

	if p.Distance(q) != 5 {
		t.Fail()
	}
}

func TestPathDistanct(t *testing.T) {
	path := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	if path.Distance() != 12 {
		t.Fail()
	}
}
