/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/simple-array-sum/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/08/hackerrank-solution-of-simple-array-sum-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=lUftBPb2gx0

*/

package main

import "fmt"

func main() {
	var n, temp, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		sum += temp
	}
	fmt.Println(sum)
}
