/*
	Find a web site that produces a large amount of data. Investigate caching by running fetch-all twice
	in succession to see whether the reported time changes much. Do you get the same content each time?
	Modify `fetchall` to print its output to a file so it can be examined.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(addr string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(addr)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	u, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("%s.html", strings.Replace(u.Path, "/", "", -1)))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", addr, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, addr)
}