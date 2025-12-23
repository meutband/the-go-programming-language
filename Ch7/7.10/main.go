// This function uses sort interface to find palindromes
package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {

	i := 0
	j := s.Len() - 1

	for i < j {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}

func main() {

	palInts := []int{1, 2, 3, 2, 1}
	fmt.Println("[1,2,3,2,1]", IsPalindrome(sort.IntSlice(palInts)))

	notPalInts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("[1,2,3,4,5,6,7,8,9]", IsPalindrome(sort.IntSlice(notPalInts)))

	palStrs := []string{"racecar", "civic", "level", "level", "civic", "racecar"}
	fmt.Println("[\"racecar\", \"civic\", \"level\", \"level\", \"civic\", \"racecar\"]", IsPalindrome(sort.StringSlice(palStrs)))

	notpalStrs := []string{"racecar", "civic", "level", "civic"}
	fmt.Println("[\"racecar\", \"civic\", \"level\", \"civic\"]", IsPalindrome(sort.StringSlice(notpalStrs)))
}
