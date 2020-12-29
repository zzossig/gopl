/*
	Implement these additional methods:

``` Go
    func (*IntSet) Len() int      // return the number of elements
    func (*IntSet) Remove(x int)  // remove x from the set
    func (*IntSet) Clear()        // remove all elements from the set
    func (*IntSet) Copy() *IntSet // return a copy of the set
```
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
	mySet := new(IntSet)
	mySet.Add(1)
	mySet.Add(3)
	mySet.Add(5)
	mySet.Add(64)
	mySet.Add(65)
	
	mySet.Add(128)
	mySet.Add(129)
	mySet.Add(130)
	mySet.Add(999)
	newSet := mySet.Copy()
	fmt.Println(newSet)
	newSet.Remove(64)

	fmt.Println(newSet)
	fmt.Println(mySet)
}

// Copy copy words in IntSet and return the words
func (s *IntSet) Copy() *IntSet {
	newSet := new(IntSet)
	for _, w := range s.words {
		newSet.words = append(newSet.words, w)
	}
	return newSet
}

// Clear clears the words in IntSet
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

// Remove removes the value of x in the IntSet
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.words[word]&(1<<bit) != 0 {
		s.words[word] -= (1<<bit)
	}
}

// Len returns a number of added value. The number of set bits is the length of the words
func (s *IntSet) Len() int {
	counts := 0
	for _, word := range s.words {
		for word > 0 {
			if word & 1 == 1 {
				counts++
			}
			word >>= 1
		}
	}
	return counts
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
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