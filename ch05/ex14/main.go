/*
	Use the `breadthFirst` function to explore a different structure.
	For example, you could use the course dependencies form the `topoSort` example (a directed graph),
	the file system hierarchy on your computer (a tree),
	or a list of bus or subway routes downloaded from your city government's web site (an undirected graph).
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	workerList := []string{"C:\\Users\\zzoss\\Github"}
	breadthFirst(folderWalker, workerList)
}

func folderWalker(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path)

	var folders []string
	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), "$") && !strings.HasPrefix(file.Name(), ".") {
			var folderPath string
			if !strings.HasSuffix(path, "\\") {
				folderPath = path + "\\" + file.Name()
			} else {
				folderPath = path + file.Name()
			}
			folders = append(folders, folderPath)
		}
	}

	return folders
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