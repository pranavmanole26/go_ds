package main

import "fmt"

// Question: https://takeuforward.org/data-structure/find-gcd-of-two-numbers/

func GCD(n1, n2 int) int {

	tn1 := n1
	tn2 := n2
	gcd := 1

	c := 2

	for {
		if tn1 == 0 || tn2 == 0 || c > tn1 || c > tn2 {
			break
		}
		if tn1%c == 0 && tn2%c == 0 {
			gcd = gcd * c
			tn1 = tn1 / c
			tn2 = tn2 / c
		} else {
			c++
		}
	}

	return gcd
}

// TODO: Not working
func GCD_BIN_SEARCH(n1, n2 int) int {

	gcd := 1

	smallerNum := n1

	if n1 > n2 {
		smallerNum = n2
	}

	l := 2
	r := smallerNum

	for {
		if l > r {
			break
		}
		m := (l + r) / 2

		if n1%m == 0 && n2%m == 0 {
			if gcd < m {
				gcd = m
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			l++
		}

	}

	return gcd
}

// Using Euclidean Algorithm
func GCD_E_A(n1, n2 int) int {

	for {
		if n1 == 0 || n2 == 0 {
			break
		}
		if n2 > n1 {
			n2 = n2 - n1
		} else {
			n1 = n1 - n2
		}
	}
	if n1 != 0 {
		return n1
	}
	return n2
}

func main() {
	n1 := 45
	n2 := 20
	fmt.Println(GCD(n1, n2))
	fmt.Println(GCD_E_A(n1, n2))
}
