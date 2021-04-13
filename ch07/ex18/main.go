/*
	Using the token-based decoder API, write a program that will read an arbitrary XML document and construct a tree of generic nodes that represents it.
	Nodes are of two kinds: `CharData` nodes represent text strings, and `Element` nodes represent named elements and their attributes.
	Each element node has a slice of child nodes.

	You may find the following declarations helpful.

	``` Go
			import "encoding/xml"
			type Node interface{} // CharData or *Element
			type CharData string
			type Element struct {
					Type     xml.Name
					Attr     []xml.Attr
					Children []Node
			}
	```
*/
package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if file, err := os.Open("./company_2.xml"); err == nil {
		defer file.Close()

		buf := bufio.NewReader(file)
		dec := xml.NewDecoder(buf)

		var stack []*Element
		var root Node
		for {
			tok, err := dec.Token()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
				os.Exit(1)
			}

			switch tok := tok.(type) {
			case xml.StartElement:
				e := &Element{Attr: tok.Attr, Type: tok.Name}
				if len(stack) == 0 {
					stack = append(stack, e)
					root = e
				} else {
					parent := stack[len(stack)-1]
					parent.children = append(parent.children, e)
					stack = append(stack, e)
				}
			case xml.EndElement:
				stack = stack[:len(stack)-1]
			case xml.CharData:
				parent := stack[len(stack)-1]
				parent.children = append(parent.children, CharData(tok))
			}
		}

		fmt.Println(root)
	}
}

type Node interface {
	node()
	String() string
}

type CharData string

func (c CharData) node() {}
func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	children []Node
}

func (e Element) node() {}
func (e Element) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("<%s", e.Type.Space+e.Type.Local))
	for _, attr := range e.Attr {
		sb.WriteString(fmt.Sprintf(" %s=%q", attr.Name.Space+attr.Name.Local, attr.Value))
	}
	sb.WriteString(">")
	for _, n := range e.children {
		switch n := n.(type) {
		case *Element:
			sb.WriteString(" " + n.String())
		case CharData:
			sb.WriteString(n.String())
		}
	}
	sb.WriteString(fmt.Sprintf("</%s>", e.Type.Space+e.Type.Local))
	return sb.String()
}
