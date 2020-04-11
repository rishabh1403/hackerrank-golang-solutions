package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)

	w := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&w[i])
	}

	for t > 0 {
		i, j := 0, 0
		fmt.Scan(&i, &j)
		min := w[i]
		for i <= j {
			if w[i] < min {
				min = w[i]
			}
			i++
		}

		fmt.Println(min)
		t--
	}
}
