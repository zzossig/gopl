/*
	Write a program that reads a single expression from the standard input, prompts the user to provide values for any variables, then evaluates the expression in the resulting environment. Handle all errors gracefully.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/ch07/ex15/eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var expr string
	variables := make(map[eval.Var]float64)

	for {
		fmt.Print("Input a single expression: ")
		if ok := scanner.Scan(); ok {
			expr = scanner.Text()
		} else {
			fmt.Println("cannot scan")
			continue
		}

		fmt.Print("Input comma, space separated variables(ex - [key] [value], ...): ")
		if ok := scanner.Scan(); ok {
			chunk := strings.Split(scanner.Text(), ",")
			for _, v := range chunk {
				kv := strings.Split(strings.TrimSpace(v), " ")
				if len(kv) != 2 {
					fmt.Println("key, value should be separated with a space")
					continue
				}
				f, err := strconv.ParseFloat(kv[1], 64)
				if err != nil {
					fmt.Println("cannot convert input to float64")
					continue
				}
				variables[eval.Var(kv[0])] = f
			}
		} else {
			fmt.Println("cannot scan")
			continue
		}

		parsed, err := eval.Parse(expr)
		if err != nil {
			fmt.Println("invalid expression")
			continue
		}

		env := eval.Env{}
		for key, value := range variables {
			env[key] = value
		}
		got := fmt.Sprintf("%.6g", parsed.Eval(env))
		fmt.Printf("result: %s\n", got)
	}
}
