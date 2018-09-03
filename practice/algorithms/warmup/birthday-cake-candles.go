/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/birthday-cake-candles/problem
blog for this code :- https://rishabh1403.com/posts/coding/hackerrank/2018/09/hackerrank-solution-of-birthday-cake-candles-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=LhgqF1vzkGc

*/

package main
import "fmt"

func main(){
	var n,t,max,occ int
	fmt.Scan(&n)

	a := []int{}
	for i:=0;i<n;i++{
		fmt.Scan(&t)
		a = append(a,t)
		if t > max {
			max = t
		}
	}

	for _,num := range a{
		if num == max{
			occ++
		}
	}

	fmt.Println(occ)

}
