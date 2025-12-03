// This file prints the outline of HTML documents
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {

	id := os.Args[1]
	for _, url := range os.Args[2:] {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()

		res := ElementByID(doc, id)
		fmt.Printf("Node: %+v\n", res)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {

	var node *html.Node

	pre := func(doc *html.Node) bool {
		for _, attr := range doc.Attr {
			if attr.Key == "id" && attr.Val == id {
				node = doc
				return true
			}
		}
		return false
	}

	forEachNode(doc, pre, nil)
	return node
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {

	if pre != nil {
		if pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		fmt.Println("post not nil")
	}
}
