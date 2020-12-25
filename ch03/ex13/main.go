// Write const declarations for KB, MB, up through YB as compactly as you can.

package main

import (
	"fmt"
	"math/big"
)

// same with the book
const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	znum := big.NewFloat(YB)
	fmt.Printf("%.0f\n", znum)
}