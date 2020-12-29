/*
	The type of each word used by `IntSet` is `uint64`,
	but 64-bit arithmetic may be inefficient on a 32-bit platform.
	Modify the program to use the `uint` type,
	which is the most efficient unsigned integer type for the platform.
	Instead of dividing by 64, define a constant holding the effective `size of uint` in bits, 32 or 64.
	You can use the perhaps too-clever expression `32 << (^uint(0) >> 63) for this purpose.
*/

package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

const sizeOfUnit = 32 << (^uint(0) >> 63)

func main() {
	fmt.Println(sizeOfUnit)
	mySet := new(IntSet)
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(4)
	mySet.Clear()
	mySet.AddAll(7,8,9,10)
	fmt.Println(mySet)
	set2 := mySet.Copy()
	set2.AddAll(101,102,103)
	mySet.IntersectWith(set2)
	fmt.Println(mySet)
}

// Elems returns a slice of words
// func (s *IntSet) Elems() []uint {
// 	r := []uint{}

// 	for i, word := range s.words {
// 		if word == 0 {
// 			continue
// 		}
// 		for j := 0; j < sizeOfUnit; j++ {
// 			if word&(1<<uint(j)) != 0 {
// 				num, err :=strconv.ParseUint(fmt.Sprintf("%d", sizeOfUnit*i+j), 0, sizeOfUnit)
// 				if err != nil {
// 					panic(err)
// 				}
// 				r = append(r, num)
// 			}
// 		}
// 	}

// 	return r
// }

// AddAll allows a list of values to be added, such as `s.AddAll(1, 2, 3)`
func (s *IntSet) AddAll(numbers ...int) {
	for _, n := range numbers {
		s.Add(n)
	}
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
	s.words = []uint{}
}

// Remove removes the value of x in the IntSet
func (s *IntSet) Remove(x int) {
	word, bit := x/sizeOfUnit, uint(x%sizeOfUnit)
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
	word, bit := x/sizeOfUnit, uint(x%sizeOfUnit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/sizeOfUnit, uint(x%sizeOfUnit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
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

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < sizeOfUnit; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", sizeOfUnit*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}