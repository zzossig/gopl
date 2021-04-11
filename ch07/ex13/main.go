/*
	Add a `String` method to `Expr` to pretty-print the syntax tree.
	Check that the results, when parsed again, yield an equivalent tree.
*/

package main

import (
	"fmt"
	"log"
	"math"

	"example.com/ch07/ex13/eval"
)

func main() {
	expr, err := eval.Parse("sqrt(A / pi)")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(expr)
	got := fmt.Sprintf("%.6g", expr.Eval(eval.Env{"A": 87616, "pi": math.Pi}))
	fmt.Println(got)
}
