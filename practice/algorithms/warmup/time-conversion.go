/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/time-conversion/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/09/hackerrank-solution-of-time-conversion-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=jX8hNDV7P9o

*/
package main

import "fmt"

func main() {
	var h, m, s int
	var suffix string
	fmt.Scanf("%d:%d:%d%s", &h, &m, &s, &suffix)
	if suffix == "AM" && h == 12 {
		h = 0
	}
	if suffix == "PM" && h != 12 {
		h += 12
	}
	fmt.Printf("%02d:%02d:%02d", h, m, s)
}
