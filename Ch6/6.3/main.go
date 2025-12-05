package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds non-negatives value x to the set
func (s *IntSet) AddAll(x ...int) {
	for _, val := range x {
		s.Add(val)
	}
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith returns a new set of common values
func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) < len(t.words) {
		for i := range s.words {
			if i < len(t.words) {
				s.words[i] &= t.words[i]
			} else {
				s.words[i] = 0
			}
		}
	} else {
		for i := range t.words {
			if i < len(s.words) {
				t.words[i] &= s.words[i]
			} else {
				t.words[i] = 0
			}
		}
		*s, *t = *t, *s
	}
}

// DifferenceWith returns a subset of set t that are not in s
func (s *IntSet) DifferenceWith(t *IntSet) {
	if len(s.words) < len(t.words) {
		tmp := s.Copy()
		tmp.IntersectWith(t)
		for i := range s.words {
			s.words[i] ^= tmp.words[i]
		}
	} else {
		tmp := t.Copy()
		tmp.IntersectWith(s)
		for i := range t.words {
			t.words[i] ^= tmp.words[i]
		}
		*s, *t = *t, *s
	}
}

// SymmetricDifference returns symmetric difference of 2 sets
func (s *IntSet) SymmetricDifference(t *IntSet) {
	if len(s.words) < len(t.words) {
		for i := range t.words {
			if i < len(s.words) {
				s.words[i] ^= t.words[i]
			} else {
				s.words = append(s.words, t.words[i])
			}
		}
	} else {
		for i := range s.words {
			if i < len(t.words) {
				t.words[i] ^= s.words[i]
			} else {
				t.words = append(t.words, s.words[i])
			}
		}
		*s, *t = *t, *s
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the length of the set
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// Copy returns a new IntSet that contains all values from s
func (s *IntSet) Copy() *IntSet {
	cpy := new(IntSet)
	cpy.words = make([]uint64, len(s.words))
	copy(cpy.words, s.words)
	return cpy
}

func main() {

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String(), x.Len())

	y.Add(9)
	y.Add(42)
	// y.Add(16)
	// y.Add(64)
	fmt.Println(y.String(), y.Len())

	x1, x2, x3, x4 := x.Copy(), x.Copy(), x.Copy(), x.Copy()
	y1, y2, y3, y4 := y.Copy(), y.Copy(), y.Copy(), y.Copy()

	x1.UnionWith(y1)
	fmt.Println(x1.String(), x1.Len())

	x2.IntersectWith(y2)
	fmt.Println(x2.String(), x2.Len())

	x3.DifferenceWith(y3)
	fmt.Println(x3.String(), x3.Len())

	x4.SymmetricDifference(y4)
	fmt.Println(x4.String(), x4.Len())
}
