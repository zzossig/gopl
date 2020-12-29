/*
	Write a variadic function `ElementsByTagName` that,
	given an HTML node tree and zero or more names,
	returns all the elements that match one of those names.
	Here are two example calls:

	``` Go
			func ElementsByTagName(doc *html.Node, name ...string) []*html.Node
			images := ElementsByTagName(doc, "img")
			headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	```
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.org/")
	checkErr(err)

	doc, err := html.Parse(bufio.NewReader(resp.Body))
	checkErr(err)
	resp.Body.Close()

	headings := elementsByTagName(doc, "h1", "h2", "h3", "h4")

	for _, heading := range headings {
		fmt.Printf("%s\n", heading.Data)
	}
}

func elementsByTagName(n *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	if n.Type == html.ElementNode {
		for _, name := range names {
			if n.Data == name {
				nodes = append(nodes, n)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, elementsByTagName(c, names...)...)
	}

	return nodes
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}