/*
	Write an in-place function that squashes each run of adjacent Unicode spaces
	(see `unicode.IsSpace`) in a UTF-8-encoded `[]byte` slice into a single ASCII space.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte("a  b      d\n\t   \ras\n\td   ")
	combineSpaces(&b)
	fmt.Printf("%s\n", b)
}

func combineSpaces(b *[]byte) {
	for i := 0; i < len(*b) - 1; {
		if unicode.IsSpace(rune((*b)[i])) {
			if unicode.IsSpace(rune((*b)[i+1])) {
				(*b)[i] = ' ' // rune(int32) automatically casted to byte(uint8)
				if i + 2 < len(*b) - 1 {
					*b = append((*b)[:i+1], (*b)[i+2:]...)
				} else {
					*b = (*b)[:i+1]
				}
			} else {
				i++
			}
		} else {
			i++
		}
	}
}