// This file inserts commas into decimal strings
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {

		s := os.Args[i]
		var ind int
		var sign byte
		if s[0] == '-' || s[0] == '+' {
			sign = s[0]
			ind++
			s = s[1:]
		}
		var dec string
		for i := ind; i < len(s); i++ {
			if s[i] == '.' {
				dec = s[i:]
				s = s[:i]
				break
			}
		}

		fmt.Println(string(sign) + comma(s) + dec)
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
