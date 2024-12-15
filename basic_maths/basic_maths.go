package main

import (
	"math"
)

// Brut force
func CountDigits(number int) (digits int) {
	c := 0
	for {
		if number != 0 {
			c++
			number = number / 10
		} else {
			break
		}
	}
	return c
}

// Optimal
func CountDigits_1(number int) (digits int) {
	// This is because, consider an example 10*10=100 meaning any number which is greater than 100 and less than 1000 will have log of n (e.g. 123) to the base 10 as 2
	// So, 2 + 1 is the number of digits it contains.
	// This is a great hack
	return int(math.Log10(float64(number)) + 1)
}

func ReverseNumber(number int) (revNumber int) {
	for {
		revNumber = revNumber*10 + number%10
		number = number / 10
		if number == 0 {
			break
		}
	}
	return revNumber
}

// Complicated
func CheckPalindrome(number int) bool {
	// create a slice and get %10 incrementally and push it to slice till half of the number size.
	// once reached half size pop numbers and compare it with the subsequent %10 of remaining number one by one.
	digits := int(math.Log10(float64(number)) + 1)
	// fmt.Print(digits)
	numSlice := make([]int, digits/2)
	popIndex := 0
	for i := 0; i < digits; i++ {
		if digits%2 == 0 {
			if i <= digits/2-1 {
				// push numbers
				numSlice[i] = number % 10
				popIndex = i
			} else {
				if numSlice[popIndex] == number%10 {
					popIndex--
				} else {
					return false
				}
			}
			number = number / 10
		} else {
			if i < digits/2 {
				// push numbers
				numSlice[i] = number % 10
				popIndex = i
			} else if i == digits/2 {
				// Skip
			} else {
				if numSlice[popIndex] == number%10 {
					popIndex--
				} else {
					return false
				}
			}
			number = number / 10
		}
	}
	return true
}

// This is a optimal approach to solve the problem
func CheckPalindrome_optimal(number int) bool {
	revNumber := 0
	tempNumber := number
	for {
		revNumber = revNumber*10 + tempNumber%10
		tempNumber = tempNumber / 10
		if tempNumber == 0 {
			break
		}
	}
	return number == revNumber
}

func main() {
	// fmt.Println(CountDigits(123))
	// fmt.Println(CountDigits_1(123))
	// fmt.Println(ReverseNumber(12300))
	// fmt.Println(CheckPalindrome(12321))
	// fmt.Println(CheckPalindrome(12322))
	// fmt.Println(CheckPalindrome(123321))
	// fmt.Println(CheckPalindrome(123331))
	// fmt.Println(CheckPalindrome(1))
	// fmt.Println(CheckPalindrome_optimal(12321))
	// fmt.Println(CheckPalindrome_optimal(12322))
	// fmt.Println(CheckPalindrome_optimal(123321))
	// fmt.Println(CheckPalindrome_optimal(123331))
	// fmt.Println(CheckPalindrome_optimal(1))
}
