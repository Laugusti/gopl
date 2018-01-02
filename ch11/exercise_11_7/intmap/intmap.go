package intmap

import (
	"fmt"
	"sort"
)

type IntMap map[int]bool

// Has reports whether the set contains the non-negative value x.
func (s *IntMap) Has(x int) bool {
	_, exists := (*s)[x]
	return exists
}

//Add adds the non-negative value x to the set.
func (s *IntMap) Add(x int) {
	(*s)[x] = true
}

// AddAll adds all the non-negative values to the set.
func (s *IntMap) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntMap) UnionWith(t *IntMap) {
	for x := range *t {
		(*s)[x] = true
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntMap) IntersectWith(t *IntMap) {
	for x := range *s {
		if !(*t)[x] {
			delete(*s, x)
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntMap) DifferenceWith(t *IntMap) {
	for x := range *t {
		s.Remove(x)
	}
}

// SymmetricDifference sets s to the elements present in once set or
// the other but not both.
func (s *IntMap) SymmetricDifference(t *IntMap) {
	s.DifferenceWith(t)
	for x, _ := range *s {
		if (*t)[x] {
			s.Remove(x)
		}
	}
}

// Len returns the number of elements.
func (s *IntMap) Len() int {
	return len(*s)
}

// Remove removes x from the set.
func (s *IntMap) Remove(x int) {
	delete(*s, x)
}

// Clear removes all elements form the set.
func (s *IntMap) Clear() {
	*s = make(IntMap)
}

// Copy returns a copy of the set
func (s *IntMap) Copy() *IntMap {
	copy := make(IntMap)
	copy.UnionWith(s)
	return &copy
}

// Elems returns a slice containing the elements of the set.
func (s *IntMap) Elems() (elements []int) {
	for x, _ := range *s {
		elements = append(elements, x)
	}
	sort.Ints(elements)
	return
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntMap) String() string {
	return fmt.Sprintf("%v", s.Elems())
}
