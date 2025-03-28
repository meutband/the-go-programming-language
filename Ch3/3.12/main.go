// This file determines in 2 inputs are anagrams of each other
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("entered incorrect number of strings, need 2")
	}
	fmt.Printf("Anagrams of '%s' and '%s'? => %t\n", os.Args[1], os.Args[2], isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int)
	for _, v := range s1 {
		m1[v]++
	}

	m2 := make(map[rune]int)
	for _, v := range s2 {
		m2[v]++
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
