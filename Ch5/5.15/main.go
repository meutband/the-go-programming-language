// This file demonstrates different mathematic vardiac functions
package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	if len(vals) == 0 {
		panic("no inputs for max, at least 1 value is required")
	}
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func min(vals ...int) int {
	if len(vals) == 0 {
		panic("no inputs for min, at least 1 value is required")
	}
	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {

	fmt.Println("sum() =", sum())
	fmt.Println("sum(1,2,3,4) =", sum(1, 2, 3, 4))

	fmt.Println("max(1,2,3) =", max(1, 2, 3))
	fmt.Println("min(1,2,3) =", min(1, 2, 3))

	// this will panic
	// fmt.Println("max() =", max())
}
