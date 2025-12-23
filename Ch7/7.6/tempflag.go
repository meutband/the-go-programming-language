// This file creates app that converts input into different degree Units
package main

import (
	"flag"
	"fmt"

	"gobook/Ch7/7.6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
