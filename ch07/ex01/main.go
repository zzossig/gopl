/*
	Using the ideas from `ByteCounter`, implement counters for words and for lines.
	You will find `bufio.ScanWords` useful.
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// Counter counts bits, words, lines
// type Counter interface {
// 	Count() int
// }

// ByteCounter represents number of bytes
type ByteCounter int

// WordCounter represents number of words
type WordCounter int

// LineCounter represents number of lines
type LineCounter int

func main() {
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"

	wc := new(WordCounter)
	lc := new(LineCounter)

	fmt.Fprintf(wc, "%s\n", input)
	fmt.Fprintf(lc, "%s\n", input)

	fmt.Println(*wc,*lc)
}

func (b *ByteCounter) Write(n []byte) (int, error) {
	*b = ByteCounter(len(n))
	return len(n), nil
}

func (w *WordCounter) Write(n []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(n))
	scanner.Split(bufio.ScanWords)
	
	counts := 0
	for scanner.Scan() {
		counts++
	}

	if err := scanner.Err(); err != nil {
		return counts, err
	}

	*w = WordCounter(counts)
	return counts, nil
}

func (l *LineCounter) Write(n []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(n))
	scanner.Split(bufio.ScanLines)

	counts := 0
	for scanner.Scan() {
		counts++
	}

	if err := scanner.Err(); err != nil {
		return counts, err
	}

	*l = LineCounter(counts)
	return counts, nil
}