// This functions prints the SHA256, SHA384, or SHA512 hashes of some input
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {

	hash := flag.String("hash", "sha256", "hash formula")
	flag.Parse()

	vals := flag.Args()
	for _, v := range vals {
		fmt.Printf("%s => %x\n", v, convert(v, hash))
	}

}

func convert(val string, hash *string) []byte {
	switch *hash {
	case "sha384":
		h := sha512.Sum384([]byte(val))
		return h[:]
	case "sha512":
		h := sha512.Sum512([]byte(val))
		return h[:]
	default:
		h := sha256.Sum256([]byte(val))
		return h[:]
	}
}
