/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/compare-the-triplets/problem
blog for this code :- https://rishabh1403.com/posts/coding/hackerrank/2018/08/hackerrank-solution-of-compare-the-triplets-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=4y8ceyunvOQ

*/

package main

import "fmt"

func main() {
	var a [3]int
	var b [3]int
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])
	var ar, br int
	for i := 0; i < 3; i++ {
		//compare here
		if a[i] > b[i] {
			ar++
		} else if b[i] > a[i] {
			br++
		}
	}
	fmt.Println(ar, br)
}
