// This file counts the number of bits that are different in 2 SHA256 hashes
package main

import (
	"crypto/sha256"
	"fmt"
	popcount "gobook/Ch2/2.3"
)

func main() {
	h1 := sha256.Sum256([]byte("hello"))
	h2 := sha256.Sum256([]byte("world"))

	fmt.Printf("%x\n", h1)
	fmt.Printf("%x\n", h2)

	fmt.Println("Number of different bits: ", diffBits(h1, h2))
}

func diffBits(h1, h2 [32]byte) int {
	var diffs int
	for i := 0; i < len(h1); i++ {
		d := h1[i] ^ h2[i]
		diffs += popcount.PopCount(uint64(d))
	}
	return diffs
}
