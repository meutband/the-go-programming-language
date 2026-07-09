// This file saves the contents of a URL into a local file.
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type response struct {
	filename string
	n        int64
	err      error
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string, ctx context.Context) (filename string, n int64, err error) {

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	resps := make(chan response)
	for _, url := range os.Args[1:] {

		go func() {
			local, n, err := fetch(url, ctx)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
				return
			}
			cancel()
			resps <- response{local, n, err}
		}()
		resp := <-resps
		fmt.Printf("%s saved => %s (%d bytes).\n", url, resp.filename, resp.n)
	}
}
