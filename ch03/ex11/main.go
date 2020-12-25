// Enhance `comma` so that it deals correctly with floating-point numbers and an optional sign.

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	if len(s) == 0 {
		return s
	}

	var sign byte
	var floating string
	startIndex := 0
	endIndex := len(s)

	if s[0] == '+' || s[0] == '-' {
		sign = s[0]
		startIndex = 1
	}
	if idx := strings.IndexRune(s, '.'); idx != -1 {
		floating = s[idx:]
		endIndex = idx
	}

	var result bytes.Buffer
	var intBuf bytes.Buffer
	intNumber := s[startIndex:endIndex]
	
	result.WriteByte(sign)
	
	for i, c := range reverse(intNumber) {
		if i != 0 && i%3==0 {
			intBuf.WriteString(",")
		}
		intBuf.WriteRune(c)
	}

	result.WriteString(reverse(intBuf.String()))
	result.WriteString(floating)
	return result.String()
}

func reverse(s string) string {
	var b bytes.Buffer
	for i := len(s)-1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}