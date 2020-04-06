/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/diagonal-difference/problem
blog for this code :- https://rishabh1403.com/posts/coding/hackerrank/2018/08/hackerrank-solution-of-diagonal-difference-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=4b0RpyRfKY8

*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)

	a := make([][]int, n)

	lsum, rsum := 0, 0

	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
			if i == j {
				lsum += a[i][j]
			}
			if i+j == n-1 {
				rsum += a[i][j]
			}
		}
	}

	diff := math.Abs(float64(lsum - rsum))
	fmt.Println(diff)

}
