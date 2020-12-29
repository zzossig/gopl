/*
	Write variadic functions `max` and `min`, analogous to `sum`.
	What should these functions do when called with no arguments?
	Write variants that require at least one argument.
*/

package main

import "fmt"


const (
	maxInt64  = 1<<63 - 1
	minInt64  = -1 << 63
)

func main() {
	fmt.Println(max2(51, 77, 743))
}

func min2(arg int, vals ...int) int {
	min := arg
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max2(arg int, vals ...int) int {
	max := arg
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	min := maxInt64
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max(vals ...int) int {
	max := minInt64
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}