/*
	Modify `forEachNode` so that the `pre` and `post` functions
	return a boolean result indicating whether to continue the traversal.
	Use it to write a function `ElementByID` with the following signature
	that finds the first HTML element with the specified `id` attribute.
	The function should stop the traversal as soon as a match is found.

	ElementByID = func(n *html.Node, id string) *html.Node {...}
*/

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

	var prev = func(n *html.Node, id string) bool {
		if elem := ElementByID(n, id); elem != nil {
			return false
		}
		return true
	}

	id := "page"
	forEachNode(doc, id, prev, nil)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) {
	if pre(n, id) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, id, pre, post)
		}
	}
	
	if post != nil {
		post(n, id)
	}
}

// ElementByID finds id in node.
func ElementByID(n *html.Node, id string) *html.Node {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			fmt.Printf("%s\t%s\n", a.Key, a.Val)
			if strings.EqualFold("id", a.Key) && strings.EqualFold(id, a.Val) {
				fmt.Println("=======================")
				fmt.Println("key found")
				fmt.Println("=======================")
				return n
			}
		}
	}
	return nil
}