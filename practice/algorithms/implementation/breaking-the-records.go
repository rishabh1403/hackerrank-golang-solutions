/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/10/hackerrank-solution-of-breaking-the-records-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=7RKyPKN1rP0

*/

package main
import (
	"fmt"
)

func main(){
	var n,min,max,cmin,cmax int
	fmt.Scan(&n)

	for i:=0;i<n;i++{
		var t int
		fmt.Scan(&t)
		if i ==0 {
			min = t
			max = t
		}else{
			if t < min {
				cmin++
				min = t
			}
			if t > max {
				cmax++
				max = t
			}
		}
	}
	fmt.Println(cmax,cmin)
}
