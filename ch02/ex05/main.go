/*
	The expression`x&(x-1)` clears the rightmost no-zero bit of `x`.
	Write a version of `PopCount` that counts bits by using this fact, and assess its performance.
*/

package main

import (
	"fmt"
	"time"
)


func main() {
	start := time.Now()
	fmt.Println(PopCountByClear(1234567891234567890))
	fmt.Printf("elapsed: %f\n", time.Since(start).Seconds())
}

// PopCountByClear count bits by clearing
func PopCountByClear(x int) int {
	counts := 0
	for x > 0 {
		x = x&(x-1)
		counts++
	}
	return counts
}