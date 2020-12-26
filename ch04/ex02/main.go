// Write a program that prints the SHA256 hash of its standard input by default but supports a command-line flag to print the SHA384 or SHA512 hash instead.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

var flag384 bool
var flag512 bool

func init() {
	flag.BoolVar(&flag384, "s3", false, "SHA384 hash")
	flag.BoolVar(&flag512, "s5", false, "SHA512 hash")
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("No input string to hash")
		return
	}
	if flag384 && flag512 {
		fmt.Println("Cannot set to true both")
		return
	}

	inputStr := strings.Join(flag.Args(), " ")

	if flag512 {
		sum := sha512.Sum512([]byte(inputStr))
		fmt.Printf("%x", sum)
	} else if flag384 {
		sum := sha512.Sum384([]byte(inputStr))
		fmt.Printf("%x", sum)
	} else {
		sum := sha256.Sum256([]byte(inputStr))
		fmt.Printf("%x", sum)
	}

}