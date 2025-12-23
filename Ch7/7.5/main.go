// This file creates a custom reader to cut data in the reader
package main

import (
	"fmt"
	"io"
	"strings"
)

type LimitedReader struct {
	r io.Reader
	n int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r: r, n: n}
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {

	if l.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.n {
		p = p[:l.n]
	}

	n, err = l.r.Read(p)
	l.n -= int64(n)
	return
}

func main() {

	demo := "this is a test message"
	l := LimitReader(strings.NewReader(demo), 7)
	blob, err := io.ReadAll(l)
	if err != io.EOF && err != nil {
		panic(err)
	}
	fmt.Println("final:", string(blob))
}
