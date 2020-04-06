/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/the-birthday-bar/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/10/hackerrank-solution-of-birthday-chocolate-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=W55rjPR7TtM

*/

package main

import (
	"fmt"
)

func main() {
	var n, d, m, c int
	fmt.Scan(&n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&d, &m)

	for i := 0; i < n; i++ {
		length, sum := 0, 0
		for j := i; j < n; j++ {
			length++
			sum += a[j]
			if sum == d && length == m {
				c++
				break
			}
			if sum > d || length > m {
				break
			}
		}
	}
	fmt.Println(c)

}
