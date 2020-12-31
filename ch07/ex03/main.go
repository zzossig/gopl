/*
	Write a `String` method for the `*tree` type in *gopl.io/ch4/treesort (§4.4)* that reveals the sequence of values in the tree.
*/

package main

import (
	"fmt"
	"strings"
)


type tree struct {
	value       string
	left, right *tree
}

func main() {
	values := []string{"love", "chicken", "football", "country", "face", "think", "rome", "time"}
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	fmt.Println(root)
}

func (t *tree) String() string {
	sb := new(strings.Builder)
	appendStrings(sb, t)
	return sb.String()
}

func appendStrings(sb *strings.Builder, t *tree) {
	if t != nil {
		appendStrings(sb, t.left)
		sb.WriteString(t.value + "--")
		appendStrings(sb, t.right)
	}
}

// Sort sorts slice of strings
func Sort(values []string) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []string, t *tree) []string {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value string) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if strings.Compare(t.value, value) > 0 {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}