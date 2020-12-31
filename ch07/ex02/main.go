/*
	Write a function `CountingWriter` with the signature below that,
	given an `io.Writer`, returns a new Writer that wraps the original,
	and a pointer to an `int64` variable that at any moment contains the number of bytes written to the new Writer.
	``` Go
    func CountingWriter(w io.Writer) (io.Writer, *int64)
	```
*/

package main

import (
	"fmt"
	"io"
)

type myWriter struct {
	l int64
	io.Writer
}

func main() {
	var w io.Writer
	w, i := CountingWriter(w)
	fmt.Fprintf(w, "%s", "abab")
	fmt.Println(*i)
	fmt.Fprintf(w, "%s", "cdcdcd")
	fmt.Println(*i)
}

// CountingWriter is an wrapper function that warps io.Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	myw := new(myWriter)
	return myw, &myw.l
}

func (mw *myWriter) Write(p []byte) (int, error) {
	mw.l = int64(len(p))
	return len(p), nil
}
