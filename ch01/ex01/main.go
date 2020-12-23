// Modify the `echo` program to also print `os.Args[0]`, the name of the command that invoked it.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, v := range os.Args {
		s += sep + v
		sep = " "
	}

	fmt.Println(s)
}