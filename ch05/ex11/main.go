/*
	The instructor of the linear algebra course decides that calculus is now a prerequisite.
	Extend the `topoSort` function to report cycles.
*/

package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"linear algebra": {"calculus"},

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

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(parent string, items []string)

	visitAll = func(parent string, items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(item, m[item])

				var cycledWith string 
				if parent != "" {
					childs := m[item]
					var achilds []string
					for _, c := range childs {
						achilds = append(achilds, m[c]...)
					}

					for _, c := range achilds {
						if c == item {
							cycledWith = parent
						}
					}
				}

				if cycledWith != "" {
					order = append(order, item + " is cycled with - " + cycledWith)
				} else {
					order = append(order, item)
				}
				cycledWith = ""
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll("", keys)
	return order
}