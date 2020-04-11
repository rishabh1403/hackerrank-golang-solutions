package main

import "fmt"

func main() {
	var year int
	fmt.Scan(&year)

	if year < 1918 {
		if year%4 == 0 {
			fmt.Printf("12.09.%v", year)
		} else {
			fmt.Printf("13.09.%v", year)
		}
	} else if year > 1918 {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			fmt.Printf("12.09.%v", year)
		} else {
			fmt.Printf("13.09.%v", year)
		}
	} else {
		fmt.Printf("26.09.1918")
	}
}
