// This file reverses a array of ints in place
package main

import "fmt"

func main() {
	arr := []byte("This is a string!")
	fmt.Println("Init:", string(arr))
	fmt.Println("Reverse:", string(reverse(arr)))
}

func reverse(arr []byte) []byte {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
