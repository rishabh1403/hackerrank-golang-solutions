/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/10/hackerrank-solution-of-divisible-sum-pairs-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=UgTCStBlcj4

*/
package main

import "fmt"

func main(){
	var n,k int
	fmt.Scan(&n,&k)

	a := make([]int,n)

	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}

	c := 0

	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if (a[i]+a[j]) % k == 0{
				c++
			}
		}
	}
	fmt.Println(c)
}
