/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/a-very-big-sum/problem
blog for this code :- https://rishabh1403.com/posts/coding/hackerrank/2018/08/hackerrank-solution-of-a-very-big-sum-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=LhgqF1vzkGc

*/

package main

import "fmt"

func main() {
	var n, temp, sum int64
	fmt.Scan(&n)
	for i := int64(0); i < n; i++ {
		fmt.Scan(&temp)
		sum += temp
	}
	fmt.Println(sum)
}
