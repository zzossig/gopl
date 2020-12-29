/*
	`(*IntSet).UnionWith` computes the union of two sets using `|`, the word-parallel bitwise OR operator.
	Implement methods for `IntersectWith`, `DifferenceWith` and `SymmetricDifference` for the corresponding set operations.
	(The symmetric difference of two sets contains the elements present in one set or the other but not both.)
*/

package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

func main() {
	set1 := new(IntSet)
	set2 := new(IntSet)
	set1.AddAll(1,2,3,4,5,6,7,8)
	set2.AddAll(6,7,8,9,10,11,12345678)
	set1.SymmetricDifference(set2)
	fmt.Println(set1)
}

// SymmetricDifference sets s to the symmdiff.. of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for len(t.words) >= len(s.words) {
		s.words = append(s.words, 0)
	}

	for i, tword := range t.words {
		s.words[i] = s.words[i] ^ tword
	}
}

// DifferenceWith sets s to the diff of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = s.words[i] ^ (s.words[i] & tword)
		}
	}
}

// IntersectWith sets s to the intersect of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
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

// AddAll allows a list of values to be added, such as `s.AddAll(1, 2, 3)`
func (s *IntSet) AddAll(numbers ...int) {
	for _, n := range numbers {
		s.Add(n)
	}
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// String returns the set as a string of the form "{1 2 3}".
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