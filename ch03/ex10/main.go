/*
	Write a non-recursive version of `comma`, using `bytes.Buffer` instead of string concatenation.
*/

package main

import (
	"bytes"
	"fmt"
	"os"
)



func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var b bytes.Buffer
	
	for i, c := range reverse(s) {
		if i != 0 && i%3==0 {
			b.WriteString(",")
		}
		b.WriteRune(c)
	}
	return reverse(b.String())
}

func reverse(s string) string {
	var b bytes.Buffer
	for i := len(s)-1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}