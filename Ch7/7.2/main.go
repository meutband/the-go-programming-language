// this file writes and counts bytes to a custom writer
package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	counter int64
	writer  io.Writer
}

func (c *CountWriter) Write(p []byte) (int, error) {
	c.counter += int64(len(p))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := CountWriter{counter: 0, writer: w}
	return &c, &c.counter
}

func main() {

	writer, count := CountingWriter(os.Stdout)
	n, err := writer.Write([]byte("This is the first sentence"))
	if err != nil {
		panic(err)
	}
	fmt.Println("count 1: ", *count, n)

	n, err = writer.Write([]byte("This is the second sentence"))
	if err != nil {
		panic(err)
	}
	fmt.Println("count 2: ", *count, n)
}
