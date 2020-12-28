// Change the `findlinks` program to traverse the `n.FirstChild` linked list using recursive calls to `visit` instead of a loop.

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.org/")
	if err != nil {
		log.Fatal(err)
	}
	
	buf := bytes.NewBuffer([]byte(""))
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	
	doc, err := html.Parse(buf)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for i, link := range visit(nil, doc) {
		fmt.Println(i, "\t", link)
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

	if c := n.FirstChild; c != nil {
		links = visit(links, c)
	}

	if s := n.NextSibling; s != nil {
		links = visit(links, s)
	}
	
	return links
}