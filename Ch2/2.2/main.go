// This file converts its numeric argument to Celcius/Fahrenheit, Feet/Meters, and Pounds/Kilograms.
package main

import (
	"fmt"
	tempconv "gobook/Ch2/2.1"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))

		ft := tempconv.Feet(t)
		m := tempconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			ft, tempconv.FToM(ft), m, tempconv.MToF(m))

		p := tempconv.Pound(t)
		kg := tempconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n",
			p, tempconv.PToKg(p), kg, tempconv.KgToP(kg))
	}
}
