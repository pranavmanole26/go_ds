package main

import "fmt"

func main() {
	num := 1
	fmt.Println(getSquareRoot(num))
}

func getSquareRoot(num int) int {
	l := 1
	r := num / 2
	sqRt := 1

	for {
		if l > r {
			return sqRt
		}
		m := (l + r) / 2
		// This condition will return nearest value of square root
		if m*m <= num {
			sqRt = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
}
