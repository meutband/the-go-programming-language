// This file prints elements from HTML documents from variadic function
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		extract(url)
	}
}

func extract(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	imgs := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	fmt.Printf("Elements - images: %d, headings: %d\n", len(imgs), len(headings))
	return nil
}

func ElementsByTagName(doc *html.Node, dtype ...string) []*html.Node {

	fmt.Println(dtype)
	var elements []*html.Node

	visitNode := func(n *html.Node) {
		for _, dt := range dtype {
			if n.Type == html.ElementNode && n.Data == dt && doc.FirstChild == nil {
				elements = append(elements, doc.FirstChild)
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return elements
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
