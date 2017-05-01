package intsetmap

import (
	"fmt"
	"sort"
	"strings"
)

// An IntMap is a set of small non-negative integers.
// Its zero value represents the empty set
type IntSet map[int]struct{}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	_, ok := (*s)[x]
	return ok
}

//Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	if *s == nil {
		*s = map[int]struct{}{}
	}
	(*s)[x] = struct{}{}
}

// AddAll adds all the non-negative values to the set.
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for k := range *t {
		s.Add(k)
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for k := range *t {
		if !s.Has(k) {
			delete(*s, k)
		}
	}
	for k := range *s {
		if !t.Has(k) {
			delete(*s, k)
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for k := range *t {
		if s.Has(k) {
			delete(*s, k)
		}
	}
}

// SymmetricDifference sets s to the elements present in one set or
// the other but not both.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for k := range *t {
		// delete key from s if it is also in t
		if s.Has(k) {
			delete(*s, k)
		} else { // add key from t to s if it isn't already in s
			s.Add(k)
		}
	}
}

// Len returns the number of elements.
func (s *IntSet) Len() int {
	return len(*s)
}

// Remove removes x from the set.
func (s *IntSet) Remove(x int) {
	delete(*s, x)
}

// Clear removes all elements form the set.
func (s *IntSet) Clear() {
	*s = nil
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var copy IntSet
	copy.UnionWith(s)
	return &copy
}

// Elems returns a slice containing the elements of the set.
func (s *IntSet) Elems() (elements []int) {
	var vals []int
	for k := range *s {
		vals = append(vals, k)
	}
	sort.Sort(sort.IntSlice(vals))
	return vals
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	vals := s.Elems()
	str := fmt.Sprintf("%v", vals)
	return fmt.Sprintf("{%s}", strings.Trim(str, "[]"))
}
