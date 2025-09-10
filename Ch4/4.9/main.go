// This file counts the frequency of words in a file
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fl := os.Args[1:][0]
	f, err := os.Open(fl)
	if err != nil {
		panic(err)
	}

	data := wordfreq(f)

	fmt.Printf("word\tcount\n")
	for w, c := range data {
		fmt.Printf("%q\t%d\n", w, c)
	}

}

func wordfreq(text io.Reader) map[string]int {

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	data := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		data[word]++
	}

	return data
}
