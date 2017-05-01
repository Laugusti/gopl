package intset

import (
	"testing"

	"github.com/Laugusti/gopl/ch11/exercise_11_2/intsetmap"
)

func mapEqualSet(m *intsetmap.IntSet, s *IntSet) bool {
	return m.String() == s.String()
}

func createSets(vals []int) (*IntSet, *intsetmap.IntSet) {
	var s IntSet
	var m intsetmap.IntSet
	for _, val := range vals {
		s.Add(val)
		m.Add(val)
	}
	return &s, &m
}

func TestAdd(t *testing.T) {
	var s IntSet
	var m intsetmap.IntSet
	if !mapEqualSet(&m, &s) {
		t.Errorf("expected: %s, got %s", &m, &s)
	}
	s.Add(100)
	m.Add(100)
	if !mapEqualSet(&m, &s) {
		t.Errorf("expected: %s, got %s", &m, &s)
	}
	s.Add(1000000)
	m.Add(1000000)
	if !mapEqualSet(&m, &s) {
		t.Errorf("expected: %s, got %s", &m, &s)
	}
}

func TestUnionWith(t *testing.T) {
	var tests = []struct {
		set1, set2 []int
	}{
		{[]int{}, []int{}},
		{[]int{}, []int{1, 2, 3}},
		{[]int{1, 100, 1000000}, []int{2, 3, 5}},
	}
	for _, test := range tests {
		s1, m1 := createSets(test.set1)
		s2, m2 := createSets(test.set2)
		s1.UnionWith(s2)
		m1.UnionWith(m2)
		if !mapEqualSet(m1, s1) {
			t.Errorf("expected: %s, got %s", m1, m2)
		}
	}
}
