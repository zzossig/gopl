/*
	Write a function that reports whether two strings are anagrams of each other,
	that is, they contain the same letter in different order.

	An anagram of a string is another string that contains the same characters, only the order of characters can be different. For example, “abcd” and “dabc” are an anagram of each other.
	Anagram Words
	LISTEN - SILENT
	TRIANGLE - INTEGRAL
*/

package main

import (
	"fmt"
	"os"
)



func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("usage: go run main.go <string1> <string2>")
	} else {
		ok := checkAna(args[0], args[1])
	
		if ok {
			fmt.Println("anagram")
		} else {
			fmt.Println("no anagram")
		}
	}
}

func checkAna(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else if len(s1) == 0 {
		return false
	}
	
	s1c := make(map[byte]int)
	s2c := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		s1c[s1[i]]++
		s2c[s2[i]]++
	}

	for k, v := range s1c {
		if s2c[k] != v {
			return false
		}
	}

	return true
}