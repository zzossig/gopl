/*
	Modify `crawl` to make local copies of the pages it finds, creating directories as necessary.
	Don't make copies of pages that come from a different domain.
	For example, if the original page comes from `golang.org`, save all files from there, but exclude ones from `vimeo.com`.
*/

package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

var hostname string

func main() {
	hostname = os.Args[1:][0]
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(path string) []string {
	u, err := url.Parse(path)
	checkErr(err)

	if strings.Contains(hostname, u.Hostname()) {
		filePath := strings.Replace(u.Path, u.RawQuery, "", -1)
		filePath = strings.Replace(filePath, "/", "\\", -1) // Since I'm running code windows 10
		
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		if _, err := os.Stat(wd + filePath); os.IsNotExist(err) {
			os.MkdirAll(wd + filePath, os.ModePerm)
		}

		f, err := os.Create(wd + filePath + "index.html")
		checkErr(err)

		resp, err := http.Get(path)
		checkErr(err)
		defer f.Close()
		defer resp.Body.Close()

		io.Copy(f, resp.Body)
	}

	list, err := links.Extract(path)
	checkErr(err)

	return list
}

func checkErr(err error) {
	if err != nil {
		log.Print(err)
	}
}