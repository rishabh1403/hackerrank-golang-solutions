/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/mini-max-sum/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/09/hackerrank-solution-of-mini-max-sum-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=-LlUwKzx-0k

*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	a := make([]int, 5)

	fmt.Scan(&a[0], &a[1], &a[2], &a[3], &a[4])
	sort.Ints(a)
	min := a[0] + a[1] + a[2] + a[3]
	max := a[1] + a[2] + a[3] + a[4]
	fmt.Println(min, max)
}
