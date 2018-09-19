/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/apple-and-orange/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/09/hackerrank-solution-of-staircase-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=f9hHJSxUfX4

*/
package main

import (
    "fmt"
)

func main(){
    
    var s,t,a,b,m,n int
    fmt.Scan(&s,&t)
    fmt.Scan(&a,&b)
    fmt.Scan(&m,&n)
    var ac, bc int
    for i:=0;i<m;i++{
        var temp int
        fmt.Scan(&temp)
        if a+temp >= s && a+temp <= t {
            ac++
        }
    }
    for i:=0;i<n;i++{
        var temp int
        fmt.Scan(&temp)
        if b+temp >= s && b+temp <= t {
            bc++
        }
    }
    fmt.Printf("%d\n%d",ac,bc)
}
