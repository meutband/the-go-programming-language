// This files prints nodes of a map in topological order
package main

import (
	"fmt"
	"os"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
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

func main() {

	courses, sorted := topoSort(prereqs)
	if !sorted {
		fmt.Println("retured repeated courses")
		os.Exit(1)
	}

	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)
	loop := make(map[string]bool)
	var visitAll func(items []string) bool

	visitAll = func(items []string) bool {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				loop[item] = true
				if !visitAll(m[item]) {
					return false
				}
				loop[item] = false
				order = append(order, item)
			} else if loop[item] {
				return false
			}
		}
		return true
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if !visitAll(keys) {
		return nil, false
	}
	return order, true
}
