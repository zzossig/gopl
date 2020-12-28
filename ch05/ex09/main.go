/*
	Write a function `expand(s string, f func(string) string) string` that replaces each substring `"$foo"` within `s` by the text returned by `f("foo")`.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var f = func(from string) (to string) {
		return "bar"
	}
	
	s := "$foo foo $foo foo $foo foo"

	s = expand(s, f)
	fmt.Println(s)
}

func expand(s string, f func(string) string) string {
	old := "$foo"
	return strings.ReplaceAll(s, old, f(old))
}