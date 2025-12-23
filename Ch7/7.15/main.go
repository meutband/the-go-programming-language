// This function takes an expression, parses and evaluates the expression
package main

import (
	"bufio"
	"fmt"
	eval "gobook/Ch7/7.14"
	"os"
	"strconv"
	"strings"
)

func getVars(key string) float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter value for %s: ", key)
	r, _ := reader.ReadString('\n')
	r = strings.ReplaceAll(r, "\n", "")
	value, err := strconv.ParseFloat(r, 64)
	if err != nil {
		fmt.Printf("%s is not a float value\n", r)
		return getVars(key)
	}
	return value
}

func main() {
	if len(os.Args) <= 1 {
		panic("missing exptression to evaluate")
	}

	expr, err := eval.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse: %v\n", err)
		os.Exit(1)
	}

	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		os.Exit(1)
	}

	env := make(map[eval.Var]float64)
	for k := range vars {
		env[k] = getVars(k.String())
	}

	result := expr.Eval(env)
	fmt.Printf("%s = %s\n", expr, strconv.FormatFloat(float64(result), 'f', -1, 64))

}
