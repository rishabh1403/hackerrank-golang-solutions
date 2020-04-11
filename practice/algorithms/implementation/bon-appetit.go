package main

import "fmt"

func main() {
	var n, k, t, sum, b int
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if i != k {
			sum += t
		}
	}
	fmt.Scan(&b)
	if b == sum/2 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - (sum / 2))
	}
}
