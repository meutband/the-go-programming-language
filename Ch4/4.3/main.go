// This file reverses a array of ints in place
package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 0}
	fmt.Println("Init", arr)
	reverse(&arr)
	fmt.Println("Reverse", arr)
}

func reverse(arr *[5]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
