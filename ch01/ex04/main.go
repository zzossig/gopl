// Modify `dup2` to print the names of all files in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsPerFile := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Please specify file names to the console")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			parseFiles(f, counts, countsPerFile)
			f.Close()
		}
	}

	result := []string{}
	for filename, items := range countsPerFile {
		inner: for text, count := range counts {
			if count > 1 {
				_, ok := items[text]
				if ok {
					result = append(result, filename)
					break inner;
				}
			}
		}
	}
	fmt.Println(result)
}

func parseFiles(f *os.File, counts map[string]int, countsPerFile map[string]map[string]int) {
	input := bufio.NewScanner(f)
	countsPerFile[f.Name()] = make(map[string]int)

	for input.Scan() {
		counts[input.Text()]++
		countsPerFile[f.Name()][input.Text()]++
	}
}