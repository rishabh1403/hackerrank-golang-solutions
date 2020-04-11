package main

import "fmt"

func main() {
	var (
		n int
		s string
	)
	fmt.Scan(&s, &n)
	var l int = len(s)

	times := n / l
	itr := n % l

	count := 0

	for i := 0; i < l; i++ {
		if s[i] == 'a' {
			count++
		}
	}

	count = count * times

	for i := 0; i < itr; i++ {
		if s[i] == 'a' {
			count++
		}
	}
	fmt.Println(count)
}
