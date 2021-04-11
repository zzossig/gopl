package eval

import (
	"fmt"
	"log"
	"testing"
)

func TestEval(t *testing.T) {
	expr, err := Parse("1 > 1")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(expr)
	got := fmt.Sprintf("%.6g", expr.Eval(Env{}))
	fmt.Println(got)
}
