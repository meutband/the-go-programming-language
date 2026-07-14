// Command 10.2 lists the files contained in an archive (ZIP or tar)
// named on the command line, using the generic archive.Read function.
// Support for each format is plugged in via the blank imports below;
// removing one drops support for that format without touching archive.Read.
package main

import (
	"fmt"
	"os"

	"gobook/Ch10/10.2/archive"
	_ "gobook/Ch10/10.2/archive/tar"
	_ "gobook/Ch10/10.2/archive/zip"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s archive-file\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "10.2: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	files, format, err := archive.Read(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "10.2: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("format: %s\n", format)
	for _, file := range files {
		fmt.Printf("%8d  %s\n", len(file.Body), file.Name)
	}
}
