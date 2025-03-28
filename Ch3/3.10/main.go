// This file inserts commas into non-negavtive decimal integer strings
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	i := n % 3
	if i == 0 {
		i = 3
	}
	buf.WriteString(s[:i])

	for j := i + 3; j < n; {
		buf.WriteString("," + s[i:j])
		i, j = j, j+3
	}
	if i < n-1 {
		buf.WriteString("," + s[i:n])
	}
	return buf.String()
}
