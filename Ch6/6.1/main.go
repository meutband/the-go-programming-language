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

// Remove deletes the non-negative value x to the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] -= 1 << bit
}

// Clear removes all values from the set
func (s *IntSet) Clear() {
	s.words = []uint64{}
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
	fmt.Println(y.String(), y.Len())

	x.UnionWith(&y)
	fmt.Println(x.String(), x.Len())

	fmt.Println(x.Has(9), x.Has(123))

	z := x.Copy()
	fmt.Println(z.String(), z.Len())

	x.Remove(144)
	x.Remove(1)

	fmt.Println(x.String(), x.Len())

	x.Clear()
	fmt.Println(x.String(), x.Len())

}
