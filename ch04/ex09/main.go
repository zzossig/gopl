/*
	Write a program `wordfreq` to report the frequency of each word in an input text file.
	Call `input.Split(bufio.ScanWords)` before the first call to `Scan` to break the input into words instead of lines.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordFreq()
}

func wordFreq() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	
	fmt.Printf("%v\n", counts)
}