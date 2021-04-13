/*
	Extend `xmlselect` so that elements may be selected not just by name, but by their attributes too,
	in the manner of CSS, so that, for instance, an element like `<div id="page" class="wide">` could be selected by a matching `id` or `class` as well as its name.
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
		var stack []string // stack of element names
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
				stack = append(stack, tok.Name.Local) // push
				if len(tok.Attr) > 0 {
					for _, attr := range tok.Attr {
						stack = append(stack, fmt.Sprintf("[%s:%s]", attr.Name, attr.Value))
					}
				}
			case xml.EndElement:
				stack = stack[:len(stack)-1] // pop
			case xml.CharData:
				if containsAll(stack, os.Args[1:]) {
					fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				}
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
