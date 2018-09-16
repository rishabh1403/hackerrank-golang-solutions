/*

Author :- Rishabh Jain <contact@rishabh1403.com>
Solution for :- https://www.hackerrank.com/challenges/grading/problem
blog for this code :-https://rishabh1403.com/posts/coding/hackerrank/2018/09/hackerrank-solution-of-grading-students-in-golang/
video explanation for this code :- https://www.youtube.com/watch?v=f9hHJSxUfX4

*/
package main

import (
    "fmt"
)

func main(){

    var n,t int
    fmt.Scan(&n)

    for i:=0;i<n;i++{
        fmt.Scan(&t)
        if t<38{
            fmt.Println(t)
        }else {
            if t%5 >2 {
                fmt.Println(((t/5)+1)*5)
            }else {
                fmt.Println(t)
            }
        }
    }
}
