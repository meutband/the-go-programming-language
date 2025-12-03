// This file prints text content from HTML documents
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

	printNonScript(doc)
}

func printNonScript(n *html.Node) {

	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "syle") {
		return
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	if n.FirstChild != nil {
		printNonScript(n.FirstChild)
	}

	if n.NextSibling != nil {
		printNonScript(n.NextSibling)
	}
}
