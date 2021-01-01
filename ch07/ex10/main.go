/*
	The `sort.Interface` type can be adapted to other use.
	Write a function `IsPalidrome(s sort.Interface) bool` that reports whether the sequence `s` is a palindrome,
	in other words, reversig the sequence would not change it.
	Assume that the elements at indices `i` and `j` are equal if `!s.Less(i, j) && !s.Less(j, i)`
*/

package main

import (
	"fmt"
	"sort"
)

type PalindromeChecker []byte

func (p PalindromeChecker) Len() int           { return len(p) }
func (p PalindromeChecker) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PalindromeChecker) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	fmt.Println(IsPalidrome(PalindromeChecker([]byte("1234432"))))
}

// IsPalidrome reports whether the sequence `s` is a palindrome
func IsPalidrome(s sort.Interface) bool {
	for i, j := 0, s.Len() - 1; i < j; i, j = i + 1, j - 1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
