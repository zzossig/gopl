// Implement `countWordAndImages`. (See Exercise 4.9 for word-splitting.)

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)



func main() {
	var checkErr = func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	b, err := ioutil.ReadFile("golang.org.txt")
	checkErr(err)

	r := bytes.NewReader(b)

	doc, err := html.Parse(r)
	checkErr(err)

	counts := make(map[string]int)
	for key, num := range countWordAndImages(counts, doc) {
		fmt.Printf("%-10s\t%d\n", key, num)
	}
}

func countWordAndImages(counts map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode && 
		n.Data != "style" && 
		n.Data != "script" && 
		n.Data != "meta" && 
		n.Data != "link" && 
		n.Data != "path" && 
		n.Data != "svg" && 
		n.Data != "html" {
		c := n.FirstChild
		if c != nil && c.Type == html.TextNode {
			tc := strings.TrimSpace(c.Data)
			if len(tc) > 0 {
				wordCount(counts, tc)	
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, i := range n.Attr {
			if i.Key == "src" {
				wordCount(counts, i.Val)
			}
		}
	}
	
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countWordAndImages(counts, c)
	}

	return counts
}

func wordCount(counts map[string]int, line string) {
	scanner := bufio.NewScanner(strings.NewReader(line))

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		for _, field := range fields {
			counts[field]++
		}
	}
}