/*
	Write a version of `PopCount` that counts bits by shifting its argument through 64-bit positions,
	testing the rightmost bit each time. Compare its performance to the table-lookup version.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	result := PopCount(1234567891234567890)
	fmt.Printf("result: %d,elapsed: %f", result, time.Since(start).Seconds())
}

// PopCount counts bits by shifting its argument
func PopCount(number int) int {
	counts := 0
	for number > 0 {
		counts += number & 1
		number = number >> 1
	}
	
	return counts
}