// Rotate over a single pass
package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5}
	ind := 2
	fmt.Println("Init", slc)
	fmt.Println("Rotated", rotate(slc, ind))
}

func rotate(slc []int, ind int) []int {
	if ind > len(slc) || ind < 0 {
		return slc
	}
	temp := make([]int, ind)
	copy(temp, slc[:ind])
	copy(slc, slc[ind:])
	copy(slc[len(slc)-ind:], temp)
	return slc
}
