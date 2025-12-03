// This file populates a map of elements types from HTML documents
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mapping: %v\n", err)
		os.Exit(1)
	}

	mp := mapping(doc)

	for elem, count := range mp {
		fmt.Printf("%s: %v\n", elem, count)
	}
}

func mapping(n *html.Node) map[string]int {

	mp := make(map[string]int)

	if n.Type == html.ElementNode {
		mp[n.Data]++
	}

	if n.FirstChild != nil {
		mapping(n.FirstChild)
	}

	if n.NextSibling != nil {
		mapping(n.NextSibling)
	}

	return mp
}
