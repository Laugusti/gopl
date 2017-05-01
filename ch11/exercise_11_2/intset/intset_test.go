package intset

import (
	"fmt"
	"testing"

	"github.com/Laugusti/gopl/ch11/exercise_11_2/intsetmap"
)

func sliceEquals(s1, s2 []int) bool {
	str1 := fmt.Sprintf("%v", s1)
	str2 := fmt.Sprintf("%v", s2)
	return str1 == str2
}
func mapEqualSet(m *intsetmap.IntSet, s *IntSet) bool {
	return m.String() == s.String()
}

func createSets(vals []int) (*IntSet, *intsetmap.IntSet) {
	var s IntSet
	var m intsetmap.IntSet
	s.AddAll(vals...)
	m.AddAll(vals...)
	return &s, &m
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		toAdd []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 1000, 1000000}},
	}
	for _, test := range tests {
		var s IntSet
		var m intsetmap.IntSet
		for _, add := range test.toAdd {
			s.Add(add)
			m.Add(add)
			if !mapEqualSet(&m, &s) {
				t.Errorf("Add %v - expected: %s, got %s", test.toAdd, &m, &s)
			}
		}
	}
}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		toAdd []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 1000, 1000000}},
	}
	for _, test := range tests {
		var s IntSet
		var m intsetmap.IntSet
		s.AddAll(test.toAdd...)
		m.AddAll(test.toAdd...)
		if !mapEqualSet(&m, &s) {
			t.Errorf("AddAll %v - expected: %s, got %s", test.toAdd, &m, &s)
		}
	}
}

func TestUnionWith(t *testing.T) {
	var tests = []struct {
		set1, set2 []int
	}{
		{[]int{}, []int{}},
		{[]int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 100, 1000000}, []int{2, 3, 5}},
	}
	for _, test := range tests {
		s1, m1 := createSets(test.set1)
		s2, m2 := createSets(test.set2)
		s1.UnionWith(s2)
		m1.UnionWith(m2)
		if !mapEqualSet(m1, s1) {
			t.Errorf("UnionWith of %v and %v - expected: %s, got %s",
				test.set1, test.set2, m1, s1)
		}
	}
}

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		set1, set2 []int
	}{
		{[]int{}, []int{}},
		{[]int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 100, 1000000}, []int{2, 3, 5}},
	}
	for _, test := range tests {
		s1, m1 := createSets(test.set1)
		s2, m2 := createSets(test.set2)
		s1.IntersectWith(s2)
		m1.IntersectWith(m2)
		if !mapEqualSet(m1, s1) {
			t.Errorf("IntersectWith of %v and %v - expected: %s, got %s",
				test.set1, test.set2, m1, s1)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		set1, set2 []int
	}{
		{[]int{}, []int{}},
		{[]int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 100, 1000000}, []int{2, 3, 5}},
	}
	for _, test := range tests {
		s1, m1 := createSets(test.set1)
		s2, m2 := createSets(test.set2)
		s1.DifferenceWith(s2)
		m1.DifferenceWith(m2)
		if !mapEqualSet(m1, s1) {
			t.Errorf("DifferenceWith of %v and %v - expected: %s, got %s",
				test.set1, test.set2, m1, s1)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		set1, set2 []int
	}{
		{[]int{}, []int{}},
		{[]int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 100, 1000000}, []int{2, 3, 5}},
	}
	for _, test := range tests {
		s1, m1 := createSets(test.set1)
		s2, m2 := createSets(test.set2)
		s1.SymmetricDifference(s2)
		m1.SymmetricDifference(m2)
		if !mapEqualSet(m1, s1) {
			t.Errorf("SymmetricDifference of %v and %v - expected: %s, got %s",
				test.set1, test.set2, m1, s1)
		}
	}
}

func TestLen(t *testing.T) {
	var tests = []struct {
		toAdd []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 1000, 1000000}},
	}
	var s IntSet
	var m intsetmap.IntSet
	for _, test := range tests {
		for _, add := range test.toAdd {
			s.Add(add)
			m.Add(add)
			if s.Len() != m.Len() {
				t.Errorf("Len %v - expected: %s, got %s", test.toAdd, &m, &s)
			}
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		set    []int
		remove []int
	}{
		{[]int{}, []int{}},
		{[]int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 100, 1000000}, []int{2, 3, 5}},
	}
	for _, test := range tests {
		s, m := createSets(test.set)
		for _, r := range test.remove {
			s.Remove(r)
			m.Remove(r)
			if !mapEqualSet(m, s) {
				t.Errorf("Remove %v (cur :%d) - expected: %s, got %s",
					test, r, m, s)
			}
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		set []int
	}{
		{[]int{}},
		{[]int{1, 2, 3}},
	}
	for _, test := range tests {
		s, m := createSets(test.set)
		s.Clear()
		m.Clear()
		if !mapEqualSet(m, s) {
			t.Errorf("Clear %v - expected: %s, got %s",
				test.set, m, s)
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		set []int
	}{
		{[]int{}},
		{[]int{1, 2, 3}},
	}
	for _, test := range tests {
		s, m := createSets(test.set)
		s, m = s.Copy(), m.Copy()
		if !mapEqualSet(m, s) {
			t.Errorf("Copy %v - expected: %s, got %s",
				test.set, m, s)
		}
	}
}

func TestElems(t *testing.T) {
	var tests = []struct {
		set []int
	}{
		{[]int{}},
		{[]int{1, 2, 3}},
	}
	for _, test := range tests {
		s, m := createSets(test.set)
		sElems, mElems := s.Elems(), m.Elems()
		if !sliceEquals(sElems, mElems) {
			t.Errorf("Elems %v - expected: %v, got %v",
				test.set, mElems, sElems)
		}
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		set []int
	}{
		{[]int{}},
		{[]int{1, 2, 3}},
	}
	for _, test := range tests {
		s, m := createSets(test.set)
		sStr, mStr := s.String(), m.String()
		if sStr != mStr {
			t.Errorf("String %v - expected: %q, got %q",
				test.set, mStr, sStr)
		}
	}
}
