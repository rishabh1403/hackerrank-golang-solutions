package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		h := 1
		var n int

		fmt.Scan(&n)
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				h++
			} else {
				h = h * 2
			}
		}
		fmt.Println(h)
		t--
	}

}
