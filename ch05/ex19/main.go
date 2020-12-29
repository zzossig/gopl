// Use `panic` and `recover` to write a function that contains no return statement yet returns a non-zero value.

package main

import "fmt"

func main() {
	myFunc()
}

func myFunc() {
	defer func() {
		recover()
		fmt.Println("???")
	}()
	panic("panic occured")
}