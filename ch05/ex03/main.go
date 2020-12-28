/*
	Write a function to print the contents of all text nodes. in an HTML document tree.
	Do not descend into `<script>` or `<style>` elements, since their contents are not visible in a web browser.
*/

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
	z := html.NewTokenizer(r)

	for _, text := range visit(nil, z) {
		fmt.Printf("%s\n", text)
	}
}

func visit(texts []string, z *html.Tokenizer) []string {
	counts := 0
	ignore := false
	loop: for {
		tt := z.Next()

		switch tt {
			case html.ErrorToken:
				break loop
			case html.StartTagToken:
				if t, _ := z.TagName(); bytes.Equal(t, []byte("script")) || bytes.Equal(t, []byte("style")) {
					ignore = true
				}
			case html.EndTagToken:
				if t, _ := z.TagName(); bytes.Equal(t, []byte("script")) || bytes.Equal(t, []byte("style")) {
					ignore = false
				}
			case html.TextToken:
				if text := bytes.TrimSpace(z.Text()); !ignore && len(text) != 0 {
					texts = append(texts, fmt.Sprintf("%-4d\t%s", counts, text))
					counts++
				}
		}
	}

	return texts
}