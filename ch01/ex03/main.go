/*
	Experiment to measure the difference in running time between our potentially
	inefficient versions and the one last uses `strings.Join`.
	(Section 1.6 illustrates part of the `time` package, and Section 11.4 shows
	how to write benchmark tests for systematic performance evaluation.)
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const argLength = 100000
var args []string

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	start := time.Now()
	var s string
	for _, arg := range args {
		s += " " + arg
	}
	fmt.Printf("Takes %d milliseconds\n", time.Since(start).Milliseconds())

	start = time.Now()
	strings.Join(args, " ")
	fmt.Printf("Takes %d milliseconds\n", time.Since(start).Milliseconds())
}

func init() {
	for i := 0; i < argLength; i++ {
		args = append(args, fmt.Sprintf("%c", rand.Intn(256)))
	}
}