// This file crawls the internet starting from URLS in the command line
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

var domain string

func crawl(u string) []string {

	if domain == "" {
		p, err := url.Parse(u)
		if err != nil {
			fmt.Println("url parse error", err)
			return nil
		}
		domain = p.Hostname()
	}

	downloadURL(u)

	list, err := Extract(u)
	if err != nil {
		log.Print(err)
	}

	out := list[:0]
	for _, l := range list {
		p, err := url.Parse(l)
		if err != nil {
			fmt.Println("url parse error", err)
			continue
		}
		if strings.Contains(p.Hostname(), domain) {
			out = append(out, l)
		}
	}
	return out
}

func downloadURL(u string) error {

	url, err := url.Parse(u)
	if err != nil {
		return fmt.Errorf("bad url: %s", err)
	}

	dir := url.Host
	var filename string
	if filepath.Ext(filename) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = url.Path
	}

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	resp, err := http.Get(u)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		resp.Body.Close()
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		resp.Body.Close()
		return err
	}

	resp.Body.Close()
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
