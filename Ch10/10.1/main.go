// The file reads an image from the standard input and writes
// it as a image to the standard output provided by the user.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	format := flag.String("format", "", "output image type. png or jpg")
	flag.Parse()

	if err := toJPEG(os.Stdin, os.Stdout, *format); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Printf("input image format: %s\n", kind)

	switch format {
	case "jpg", "jpeg":
		return jpeg.Encode(out, img, nil)
	case "png":
		return png.Encode(out, img)
	}

	return fmt.Errorf("invalid output image format")
}
