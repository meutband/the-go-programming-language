// This file counts the Unicode characters
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {

	counts := make(map[string]int)
	in := bufio.NewReader(os.Stdin)

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar {
			counts["invalid"]++
		}
		if unicode.IsControl(r) {
			counts["control"]++
		}
		if unicode.IsDigit(r) {
			counts["digit"]++
		}
		if unicode.IsGraphic(r) {
			counts["graphic"]++
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		}
		if unicode.IsMark(r) {
			counts["mark"]++
		}
		if unicode.IsNumber(r) {
			counts["number"]++
		}
		if unicode.IsPunct(r) {
			counts["punct"]++
		}
		if unicode.IsSpace(r) {
			counts["space"]++
		}
		if unicode.IsSymbol(r) {
			counts["symbol"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
