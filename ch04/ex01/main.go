/*
	Write a function that counts the number of bits that are different in two SHA256 hashes.
	(See PopCount from Section 2.6.2.)
*/

package main

import (
	"crypto/sha256"
	"fmt"
)


func main() {
	h1 := sha256.New()
	h2 := sha256.New()
	h1.Write([]byte("hello "))
	h2.Write([]byte("world!"))

	num := compareBits(h1.Sum(nil), h2.Sum(nil))
	fmt.Println(num)
}

func compareBits(b1 []byte, b2 []byte) byte {
	var count byte
	for i := 0; i < len(b1); i++ {
		b1b2xor := b1[i]^b2[i]
		
		for b1b2xor > 0 {
			count += b1b2xor & 1
			b1b2xor >>= 1
		}
		
	}
	return count
}