/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/between-two-sets/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/09/hackerrank-solution-of-between-two-sets-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=LJ6nUUZpBiE

*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	socks := make([]int, 101)

	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)

		socks[t]++
	}

	count := 0

	for _, v := range socks {
		count += v / 2
	}

	fmt.Println(count)
}
