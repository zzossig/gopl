/*
	Develop `startElement` and `endElement` into a general HTML pretty-printer.
	Print comment nodes, text nodes, and the attributes of each element (`<a href='...'>`).
	Use short forms like `<img/>` instead of `<img></img>` when an element has no children.
	Write a test to ensure that the output can be parsed successfully.
	(See Chapter 11.)
*/

// I didn't create startElement and endElement. Actually, the question is too ambigious to understand for me

package main

import (
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

	doc, err := html.Parse(bytes.NewReader(b))
	checkErr(err)

	depth := 0
	var sb strings.Builder

	var startElement = func(n *html.Node) {
		var attrs []string
		tagName := n.Data

		for _, v := range n.Attr {
			attrs = append(attrs, fmt.Sprintf(" %s=%q", v.Key, v.Val))
		}

		sb.WriteString(fmt.Sprintf("%*s<%s", depth * 2, "", tagName))
		for _, attr := range attrs {
			sb.WriteString(attr)
		}

		if c := n.FirstChild; c != nil {
			sb.WriteString(fmt.Sprintln(">"))
			if c.Type == html.TextNode && c.Parent.Data != "script" {
				text := strings.TrimSpace(c.Data)
				if len(text) > 0 {
					sb.WriteString(fmt.Sprintf("%*s%s\n", depth * 2 + 2 , "", text))
				}
			}
		} else {
			sb.WriteString(fmt.Sprintln(" />"))
		}
		
		depth++
	}

	var endElement = func(n *html.Node) {
		depth--
		tagName := n.Data

		if c := n.FirstChild; c != nil {
			sb.WriteString(fmt.Sprintf("%*s</%s>\n", depth * 2, "", tagName))
		}
	}

	forEachNode(doc, startElement, endElement)
	fmt.Println(sb.String())
}

func forEachNode(n *html.Node, prev, post func(*html.Node)) {
	if prev != nil {
		if n.Type == html.ElementNode {
			if n.Parent.Data != "svg" {
				prev(n)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, prev, post)
	}

	if post != nil {
		if n.Type == html.ElementNode {
			if n.Parent.Data != "svg" {
				post(n)
			}
		}
	}
}


// if n.Type == html.CommentNode {
		// 	sb.WriteString("[CommentNode]")
		// 	sb.WriteString(n.Data)
		// 	sb.WriteByte('\n')
		// }

		// if n.Type == html.TextNode {
		// 	if n.Parent.Data != "script" {
		// 		if str := strings.TrimSpace(n.Data); len(str) > 0 {
		// 			sb.WriteString("[TextNode]")
		// 			sb.WriteString(n.Data)
		// 			sb.WriteByte('\n')
		// 		}
		// 	}
		// }