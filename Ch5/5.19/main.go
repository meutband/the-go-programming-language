// This file demonstrates panic recovery
package main

import "fmt"

func divide(n, d int) int {
	var res int
	defer func() {
		if p := recover(); p != nil {
			res = 0
			fmt.Println("divide by zero")
		}
	}()
	res = n / d
	return res
}

func main() {
	fmt.Println("4/2 = ", divide(4, 2))
	fmt.Println("4/0 = ", divide(4, 0))
}
