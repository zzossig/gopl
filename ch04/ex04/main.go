// Write a version of `rotate` that operates in a single pass.

package main

import "fmt"



func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	
	rotate(s, 1)
	rotate(s, -1)
	rotate(s, 1)
	fmt.Println(s)
}

func rotate(s []int, dir int) {
	if len(s) == 0 || len(s) == 1 {
		return
	}

	if dir < 0 { // rotate left
		first := s[0]
		other := s[1:]
		for i := 0; i < len(s); i++ {
			if i == len(s) - 1 {
				s[i] = first
			} else {
				s[i] = other[i]
			}
		}
	} else { // rorate right
		last := s[len(s)-1]
		other := s[:len(s)-1]
		for i := len(s) - 1; i >= 0; i-- {
			if i == 0 {
				s[i] = last
			} else {
				s[i] = other[i-1]
			}
		}
	}
}