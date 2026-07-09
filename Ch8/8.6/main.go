// This package crawls the website to a user provided depth
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// maxDepth is the max crawl size
var maxDepth int

type link struct {
	url   string
	depth int
}

func crawl(l link) []link {
	fmt.Println(l.url)
	var ls []link
	if l.depth < maxDepth {
		depth := l.depth + 1
		list, err := links.Extract(l.url)
		if err != nil {
			log.Print(err)
		}

		for _, url := range list {
			ls = append(ls, link{url, depth})
		}
	}
	return ls
}

func main() {

	flag.IntVar(&maxDepth, "depth", 3, "max crawl depth")
	flag.Parse()

	worklist := make(chan []link)  // lists of URLs, may have duplicaties
	unseenLinks := make(chan link) // de-duplicated URLs

	//Add command-line arguments to worklist.
	go func() {
		var ls []link
		for _, url := range flag.Args() {
			ls = append(ls, link{url, 0})
		}
		worklist <- ls
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-deplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[link]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
