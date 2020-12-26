/*
	Modify `charcount` to count letters, digits, and so on in their Unicode categories,
	using functions like `unicode.IsLetter`.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	letterCount := make(map[rune]int)
	digitCount := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCount[r]++
		}
		if unicode.IsDigit(r) {
			digitCount[r]++
		}
		utflen[n]++
	}
	fmt.Printf("rune\tletter count\n")
	for c, n := range letterCount {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("rune\tdigit count\n")
	for c, n := range digitCount {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}