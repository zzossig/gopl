// Write a variadic version of `strings.Join`.

package main

import (
	"fmt"
	"strings")


func main() {
	fmt.Println(stringsJoin([]string{"a", "b", "c", "d"}, "\t"))
}

func stringsJoin(elems []string, sep string) string {
	var s strings.Builder
	for i, elem := range elems {
		s.WriteString(elem)
		if i < len(elems) {
			s.WriteString(sep)
		}
	}
	return s.String()
}