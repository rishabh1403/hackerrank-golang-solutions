package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for n > 0 {
		var a, b int
		fmt.Scan(&a, &b)

		low := math.Ceil(math.Sqrt(float64(a)))
		high := math.Floor(math.Sqrt(float64(b)))
		fmt.Printf("%.0f\n", (high + 1 - low))
		n--
	}
}
