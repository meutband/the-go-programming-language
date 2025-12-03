// This file demonstrates join function as a vardiac function
package main

import (
	"fmt"
)

func join(sep string, vals ...string) string {

	if len(vals) == 0 {
		return ""
	}
	if len(vals) == 1 {
		return vals[0]
	}

	var combined string
	for _, v := range vals[:len(vals)-1] {
		combined = combined + v + sep
	}

	combined = combined + vals[len(vals)-1]
	return combined
}

func main() {
	fmt.Println("join('%', 'hello', 'world') =", join("%", "hello", "world"))
	fmt.Println("join('^', 'hello') =", join("^", "hello"))
	fmt.Println("join('@') =", join("@"))
}
