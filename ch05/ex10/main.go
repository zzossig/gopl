/*
	Rewrite `topoSort` to use `maps` instead of `slices` and eliminate the initial sort.
	Verify that the results, though nondeterministic, are valid topological orderings.
*/

package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	courses := topoSort(prereqs)

	for i := 0; i < len(courses); i++ {
		fmt.Printf("%d\t%s\n", i, courses[i])
	}
}

func topoSort(m map[string][]string) map[int]string {
	orders := make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)

	visitAll = func(ms map[string][]string) {
		for key, items := range ms {
			if !seen[key] {
				seen[key] = true
				
				// map value doesn't matter
				nms := make(map[string][]string)
				for _, v := range items {
					nms[v] = []string{}
				}
				visitAll(nms)
				
				orders[len(orders)] = key
			}
		}
	}

	visitAll(m)

	return orders
}