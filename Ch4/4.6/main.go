// This file removes extra spaces in a string
package main

import (
	"fmt"
	"unicode"
)

func main() {
	init := []byte("this  is a\tstring!")
	cpy := make([]byte, len(init))
	copy(cpy, init)
	fmt.Printf("'%s' => '%s'\n", string(cpy), string(squash(init)))
}

func squash(input []byte) []byte {

	var inx int
	var consec bool
	for _, s := range string(input) {
		if unicode.IsSpace(s) {
			if consec {
				continue
			} else {
				consec = true
				s = ' '

			}
		} else {
			consec = false
		}
		input[inx] = byte(s)
		inx++
	}
	return input[:inx]
}
