package main

import "fmt"

func main() {
	var n int
	var s string

	fmt.Scan(&n, &s)

	depth := 0
	valley := 0
	for i := 0; i < n; i++ {
		if s[i] == 'U' {
			depth++
		}
		if s[i] == 'D' {
			depth--
		}

		if depth == 0 && s[i] == 'U' {
			valley++
		}
	}
	fmt.Println(valley)
}
