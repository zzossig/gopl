/*
	Define a new concrete type that satisfies the `Expr` interface and provides a new operation such as computing the minimum value of its operands.
	Since the `Parse` function does not create instances of this new type, to use it you will need to construct a syntax tree directly (or extend the parser).
*/

package main

import (
	"fmt"
	"log"

	"example.com/ch07/ex14/eval"
)

func main() {
	expr, err := eval.Parse("1 < 1.1")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(expr)
	got := fmt.Sprintf("%t", expr.Eval(eval.Env{}))
	fmt.Println(got)
}
