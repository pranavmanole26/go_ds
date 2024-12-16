package main

import "fmt"

func PrintOneToN(i, n int) {
	if i > n {
		return
	}
	fmt.Println(i)
	PrintOneToN(i+1, n)
}

func PrintNToOne(i, n int) {
	if i < n {
		return
	}
	fmt.Println(i)
	PrintNToOne(i-1, n)
}

func SumOfFirstNNumbers(i, n, sum int) int {
	if i > n {
		return sum
	}
	sum += i
	return SumOfFirstNNumbers(i+1, n, sum)
}

func FactorialOfNNumbers(n, factorial int) int {
	if n == 1 {
		return factorial
	}
	factorial = factorial * n
	return FactorialOfNNumbers(n-1, factorial)
}

func ReverseArray(i int, input []int) []int {
	if i >= len(input)/2 {
		return input
	}
	input[i], input[len(input)-(i+1)] = input[len(input)-(i+1)], input[i]
	return ReverseArray(i+1, input)
}

func IsPalindrome(i int, str string) bool {
	if i >= len(str)/2 {
		return true
	}
	if str[i] != str[len(str)-(i+1)] {
		return false
	}
	return IsPalindrome(i+1, str)
}

func Fibonaci(i, n int, fib []int) []int {
	if i > n {
		return fib
	}
	if i < 2 {
		fib = append(fib, i)
	} else {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return Fibonaci(i+1, n, fib)
}

func main() {
	// PrintOneToN(1, 5)
	// PrintNToOne(5, 1)
	// fmt.Println(SumOfFirstNNumbers(1, 10, 0))
	// fmt.Println(FactorialOfNNumbers(5, 1))
	// fmt.Println(ReverseArray(0, []int{1, 2, 3, 4, 5, 6}))
	// fmt.Println(ReverseArray(0, []int{1, 2, 3, 4, 5}))
	// fmt.Println(ReverseArray(0, []int{1}))
	// fmt.Println(IsPalindrome(0, "abcba"))
	// fmt.Println(IsPalindrome(0, "abacba"))
	// fmt.Println(IsPalindrome(0, "abccba"))
	fmt.Println(Fibonaci(0, 5, []int{}))
	fmt.Println(Fibonaci(0, 1, []int{}))
	fmt.Println(Fibonaci(0, 10, []int{}))
}
