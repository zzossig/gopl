// Extend the `visit` function so that it extracts other kinds of links from the document, such as image, scripts, and style sheets.

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

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

	for _, link := range visit(nil, doc) {
		fmt.Printf("%s\n", link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, i := range n.Attr {
			if i.Key == "src" {
				links = append(links, i.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "link" {
		for _, l := range n.Attr {
			if l.Key == "href" {
				links = append(links, l.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "script" {
		for _, l := range n.Attr {
			if l.Key == "src" {
				links = append(links, l.Val)
			}
		}
	}

	if c := n.FirstChild; c != nil {
		links = visit(links, c)
	}

	if s := n.NextSibling; s != nil {
		links = visit(links, s)
	}
	
	return links
}