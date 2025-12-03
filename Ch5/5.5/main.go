// This file counts the number of words and images in HTML documents
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	for _, url := range os.Args[1:] {

		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Printf("%s - CountWordsAndImages: %v\n", url, err)
			continue
		}

		fmt.Printf("%s - words: %v, images: %v\n", url, words, images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing html: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImages extracts word counts and image counts from documents
func countWordsAndImages(n *html.Node) (words, images int) {

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	if n.Type == html.TextNode {
		reader := strings.NewReader(n.Data)
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}

	if !(n.Data == "script" || n.Data == "style") {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			word, image := countWordsAndImages(c)
			words += word
			images += image
		}
	}

	return
}
