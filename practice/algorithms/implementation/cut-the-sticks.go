package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	for {
		min := 1001
		length := 0
		for i := 0; i < n; i++ {
			if nums[i] != 0 {

				length++
				if nums[i] < min {
					min = nums[i]
				}
			}

		}
		for i := 0; i < n; i++ {
			if nums[i] != 0 {
				nums[i] = nums[i] - min
			}

		}
		if length == 0 {
			break
		}
		fmt.Println(length)
	}
}
