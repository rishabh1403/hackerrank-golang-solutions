/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/drawing-book/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2020/04/hackerrank-solution-of-drawing-book-in-golang
Youtube video :- https://youtu.be/yTbd7vRDF4Y

*/

package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)
	front := p / 2

	back := (n - p) / 2

	if n%2 == 0 {
		back = (n + 1 - p) / 2
	}

	if front < back {
		fmt.Println(front)
	} else {
		fmt.Println(back)
	}
}
