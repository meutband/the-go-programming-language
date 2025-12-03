// This file prints the outline of HTML documents
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()

		forEachNode(doc, startElement, endElement)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {

	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {

	if n.Type == html.ElementNode {

		var attrs string
		for _, attr := range n.Attr {
			attrs += fmt.Sprintf("%s=%s ", attr.Key, attr.Val)
		}

		var last string
		if n.FirstChild == nil {
			last = " /"
		}

		fmt.Printf("%*s<%s%s%s>\n", depth*2, "", n.Data, attrs, last)
		if n.FirstChild != nil {
			depth++
		}
	}

	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) != 0 {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		}
	}

	if n.Type == html.CommentNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {

	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
