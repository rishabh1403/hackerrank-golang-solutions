/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/plus-minus/problem
blog for this code :- https://rishabh1403.com/posts/coding/hackerrank/2018/08/hackerrank-solution-of-plus-minus-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=7996I_bqVmk

*/

package main

import "fmt"

func main() {
	var l, p, n, z, t int
	fmt.Scan(&l)

	for i := 0; i < l; i++ {
		fmt.Scan(&t)
		if t > 0 {
			p++
		} else if t < 0 {
			n++
		} else {
			z++
		}
	}

	pf := float64(p) / float64(l)
	nf := float64(n) / float64(l)
	zf := float64(z) / float64(l)

	fmt.Printf("%.6f\n", pf)
	fmt.Printf("%.6f\n", nf)
	fmt.Printf("%.6f\n", zf)

}
