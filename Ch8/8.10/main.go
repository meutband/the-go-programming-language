// This file crawls http files until timeout is called
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

const timeout = 2 * time.Second

func crawl(url string, ctx context.Context) []string {
	fmt.Println(url)
	list, err := Extract(url, ctx)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	ctx, cancel := context.WithCancel(context.Background()) // for cancel http.Request

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for {
				select {
				case link := <-unseenLinks:
					foundLinks := crawl(link, ctx)
					go func() {
						select {
						case <-ctx.Done(): // if closed, stop foundlinks return
						default:
							worklist <- foundLinks
						}
					}()
				case <-ctx.Done(): // if closed, stop unseenlinks
					return

				}
			}
		}()
	}

	go func() {
		time.Sleep(timeout)
		log.Println("context canceled")
		cancel()
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
