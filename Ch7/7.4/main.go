// This file creates a custom reader to parse HMTL
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type StringReader struct {
	s string
}

func NewStringReader(s string) *StringReader {
	return &StringReader{s: s}
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

// from section 5.2
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {

	resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	io.Copy(&buf, resp.Body)

	sr := NewStringReader(buf.String())
	doc, err := html.Parse(sr)
	if err != nil {
		panic(err)
	}

	outline(nil, doc)
}
