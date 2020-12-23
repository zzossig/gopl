/*
	Rewrite `PopCount` to use a loop instead of a single expression. Compare the performance of two versions.
	(Section 11.4 shows how to compare the performance of different implementations systematically.)
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var pc [256]byte

func main() {
	const testNum = 1234567891234567890
	start := time.Now()
	result1 := PopCountLoop(testNum)
	fmt.Printf("PopCountLoop, elapsed: %f\n", time.Since(start).Seconds())

	start = time.Now()
	result2 := PopCount(testNum)
	fmt.Printf("PopCount, elapsed: %f\n", time.Since(start).Seconds())

	fmt.Println(result1, result2)
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
/*
	That is to say,
	number   x   number of set bits
	0000     0			0
	0001     0			1
	0010     0			1
	0011     0			2
	0100     0			1
	0101     0			2
	0110     0			2
	0111     0			3
	1000     0			1
	1001     0			2
	1010     0			2
	1011     0			3
	1100     0			2
	1101     0			3
	1110     0			3
	1111     0			4
	...
*/
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop is the loop version of the PopCount
func PopCountLoop(x uint64) int {
	formatted := fmt.Sprintf("%v", strconv.FormatInt(int64(x), 2))
	splitted := strings.Split(formatted, "")

	var count int
	for _, v := range splitted {
		if n, err := strconv.Atoi(v); err == nil {
			if n == 1 {
				count++
			}
		}
	}

	return count
}