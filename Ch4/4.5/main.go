// filter duplicates from slice
package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 3, 4, 4}
	fmt.Println("Init", slc)
	fmt.Println("Filtered", filter(slc))
}

func filter(slc []int) []int {
	if len(slc) == 0 {
		return slc
	}
	i := 0
	for j := 1; j < len(slc); j++ {
		if slc[i] != slc[j] {
			i++
			slc[i] = slc[j]
		}
	}
	return slc[:i+1]
}
