package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	clouds := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&clouds[i])
	}

	position, jumps := 0, 0
	for position < n-1 {
		if position != n-2 {
			if clouds[position+2] == 0 {
				position += 2
			} else {
				position++
			}
		} else {
			position++
		}
		jumps++

	}
	fmt.Println(jumps)
}
