package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for t > 0 {
		var n, c, m int

		fmt.Scan(&n, &c, &m)

		wrappers := n / c

		ans := n / c

		for wrappers >= m {
			ans = ans + (wrappers / m)
			wrappers = (wrappers / m) + (wrappers % m)
		}
		fmt.Println(ans)
		t--
	}
}
