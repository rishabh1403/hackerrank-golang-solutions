/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/migratory-birds/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/09/hackerrank-solution-of-staircase-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=f9hHJSxUfX4

*/
package main

import "fmt"

func main(){
	var n,t int

	fmt.Scan(&n)

	a := make([]int,5)

	for i:=0;i<n;i++{
		fmt.Scan(&t)
		t = t -1
		a[t]++
	}
	max,occ := a[0],0

	for i,num := range a{
		if num > max{
			max = num
			occ = i+1
		}
	}
	fmt.Println(occ)
}
