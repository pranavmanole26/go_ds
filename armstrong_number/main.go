package main

// Question: https://takeuforward.org/maths/check-if-a-number-is-armstrong-number-or-not/

import (
	"fmt"
	"math"
)

func IsArmstrongNumber(n int) bool {
	digits := getDigits(n)
	// fmt.Println(digits)
	sum := 0
	tempNum := n

	for {
		if tempNum == 0 {
			break
		}
		mod := tempNum % 10
		sum = sum + int(math.Pow(float64(mod), float64(digits)))
		tempNum = tempNum / 10
	}

	return n == sum
}

func getDigits(n int) int {
	digits := 0
	for {
		if n == 0 {
			break
		}
		digits++
		n = n / 10
	}
	return digits
}

func main() {
	fmt.Println(IsArmstrongNumber(1234))
	fmt.Println(IsArmstrongNumber(153))
	fmt.Println(IsArmstrongNumber(371))
}
