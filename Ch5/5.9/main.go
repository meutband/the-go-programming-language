// This functions alters a string with the results of a function
package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string{"$food", "$foolis", "$football"}

	for _, word := range words {
		fmt.Printf("To Upper: %s -> %s\n", word, expand(word, strings.ToUpper))
		fmt.Printf("Reverse: %s -> %s\n", word, expand(word, reverse))
	}
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
