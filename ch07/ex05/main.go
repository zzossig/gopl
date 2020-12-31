/*
	The `LimitReader` function in the io package accepts an `io.Reader r` and a number of bytes `n`,
	and returns another `Reader` that reads from `r` but reports an end-of-file condition after n bytes.
	Implement it.

	``` Go
			func LimitReader(r io.Reader, n int64) io.Reader
	```
*/

package main

import (
	"bytes"
	"fmt"
	"io"
)

type myReader struct {
	limit int64
	r io.Reader
}

func main() {
	r := bytes.NewReader([]byte("abcdefghijklmnopqrstuvwxyz12345"))
	mr := LimitReader(r, 5)
	s := make([]byte, 6)

	n, err := mr.Read(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\t%d\n", s, n)

	// check built in LimitReader what's different with my LimitReader
	rr := io.LimitReader(r, 5)
	ss := make([]byte, 7)

	nn, err := rr.Read(ss)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\t%d\n", ss, nn)
	
	// for {
	// 	if _, err := mr.Read(s) ; err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Printf("%s\n", s)
	// }
	// fmt.Println("the end")
}

// LimitReader is an io.Reader wrapper
func LimitReader(r io.Reader, n int64) io.Reader {
	var mr myReader = myReader{n, r}
	return mr
}

func (mr myReader) Read(p []byte) (int, error) {
	if len(p) > int(mr.limit) {
		pp := make([]byte, mr.limit)
		n, err := mr.r.Read(pp)

		if err != nil {
			return n, err
		}

		p = append(p[:0], pp...)
		return n, io.EOF
	}
	return mr.r.Read(p)
}