package intset

import (
	"bytes"
	"fmt"
)

const uintsize = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/uintsize, uint(x%uintsize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

//Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/uintsize, uint(x%uintsize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds all the non-negative values to the set.
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := len(s.words) - 1; i >= 0; i-- {
		if i >= len(t.words) {
			s.words = s.words[:i]
		} else {
			s.words[i] &= t.words[i]
			if s.words[i] == 0 && len(s.words) == i+1 {
				s.words = s.words[:i]
			}
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := len(s.words) - 1; i >= 0; i-- {
		if i >= len(t.words) {
			continue
		} else {
			s.words[i] &^= t.words[i]
			if s.words[i] == 0 && len(s.words) == i+1 {
				s.words = s.words[:i]
			}
		}
	}
}

// SymmetricDifference sets s to the elements present in once set or
// the other but not both.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i := len(s.words) - 1; i >= 0; i-- {
		if i >= len(t.words) {
			continue
		} else {
			s.words[i] ^= t.words[i]
			if s.words[i] == 0 && len(s.words) == i+1 {
				s.words = s.words[:i]
			}

		}
	}
}

// Len returns the number of elements.
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for bit := uint(0); bit < uintsize; bit++ {
			if word&(1<<bit) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove removes x from the set.
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	word, bit := x/uintsize, uint(x%uintsize)
	s.words[word] &^= 1 << bit
}

// Clear removes all elements form the set.
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var copy IntSet
	copy.UnionWith(s)
	return &copy
}

// Elems returns a slice containing the elements of the set.
func (s *IntSet) Elems() (elements []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintsize; j++ {
			if word&(1<<uint(j)) != 0 {
				elements = append(elements, uintsize*i+j)
			}
		}
	}
	return
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < uintsize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", uintsize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
