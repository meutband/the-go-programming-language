// This program measures difference in time of execution between our 3 functions
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	// fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func echo3(args []string) {
	// fmt.Println(strings.Join(args, " "))
	strings.Join(args[1:], " ")
}

func main() {

	args := os.Args

	t1 := time.Now()
	echo1(args)
	fmt.Println("Echo1: ", time.Since(t1))

	t2 := time.Now()
	echo2(args)
	fmt.Println("Echo2: ", time.Since(t2))

	t3 := time.Now()
	echo3(args)
	fmt.Println("Echo3: ", time.Since(t3))
}
