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

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for k := range *t {
		s.Add(k)
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var vals []int
	for k := range *s {
		vals = append(vals, k)
	}
	sort.Sort(sort.IntSlice(vals))
	str := fmt.Sprintf("%v", vals)
	return fmt.Sprintf("{%s}", strings.Trim(str, "[]"))
}
