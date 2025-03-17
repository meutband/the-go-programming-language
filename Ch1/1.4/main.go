// This file prints the text of lines and all the files that the lines appear in.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}

	for line, files := range counts {
		// only print where lines are repeated in the files
		if len(files) > 1 {
			fmt.Printf("'%s'\t%v\n", line, files)
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		counts[line] = append(counts[line], f.Name())
	}
}
