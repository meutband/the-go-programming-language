// This file reads words from bytes and counts the number of words and lines
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(bytes.NewReader(p))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		*c++
	}
	return len(p), scan.Err()
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(bytes.NewReader(p))
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		*c++
	}
	return len(p), scan.Err()
}

func main() {

	var bc ByteCounter
	_, err := bc.Write([]byte("demo byte counter"))
	if err != nil {
		panic(err)
	}
	fmt.Println("byte counter: ", bc)

	var wc WordCounter
	_, err = wc.Write([]byte("demo word counter"))
	if err != nil {
		panic(err)
	}
	fmt.Println("word counter: ", wc)

	var lc LineCounter
	_, err = lc.Write([]byte("line\n word count\ner\n!!"))
	if err != nil {
		panic(err)
	}
	fmt.Println("line counter: ", lc)
}
