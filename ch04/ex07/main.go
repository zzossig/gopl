/*
	Modify `reverse` to reverse the characters of a `[]byte` slice that represents a UTF-8-encoded string, in place.
	Can you do it without allocating new memory?
*/

package main

import "fmt"


func main() {
	b := []byte("abcdefghijklmnopqrstuvwxyz")
	reverse2(b)
	fmt.Printf("%s\n", b)
}

// memory usage: i, j, b
func reverse2(b []byte) {
	left := b[:0]
	for i := 0; i < len(b); i++ {
		other := b[i:len(b) - 1]
		j := b[len(b) - 1]
		b = append(left[:i+1], other...)
		left = append(left, j)
	}
}

// original, memory usage: i, j, b, temp(for swapping two slices)
func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}