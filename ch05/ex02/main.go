// Write a function to populate a mapping from element names--`p, div, span, and so on`--to the number of elements with that name in an HTML document tree.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/html"
)

func main() {
	countElem()
}

func countElem() {
	var checkErr = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	b, err := ioutil.ReadFile("golang.org.txt")
	checkErr(err)

	buf := bytes.NewReader(b)

	doc, err := html.Parse(buf)
	checkErr(err)

	counts := make(map[string]int)
	for key, value := range visit(counts, doc) {
		fmt.Printf("%-10s%d\n", key, value)
	}
}

func visit(counts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	if c := n.FirstChild; c != nil {
		counts = visit(counts, c)
	}

	if s := n.NextSibling; s != nil {
		counts = visit(counts, s)
	}
	
	return counts
}