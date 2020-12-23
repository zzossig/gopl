/*
	The function call `io.Copy(dst, src)` reads from `src` and writes to `dst`.
	Use it instead of `ioutil.ReadAll` to copy the response body to `os.Stdout`
	without requiring a buffer large enough to hold the entire stream.
	Be sure to check the error result of `io.Copy`.
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		buf := bytes.NewBuffer([]byte(""))
		nbytes, err := io.Copy(buf, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n %d bytes copied\n", buf.Bytes(), nbytes)
		fmt.Printf("Takes %d milliseconds\n", time.Since(start).Milliseconds())
	}
}