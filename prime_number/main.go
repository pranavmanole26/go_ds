package main

import "fmt"

func IsPrimeNumber(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrimeNumber(10))
	fmt.Println(IsPrimeNumber(2))
	fmt.Println(IsPrimeNumber(3))
	fmt.Println(IsPrimeNumber(47))
	fmt.Println(IsPrimeNumber(49))
	fmt.Println(IsPrimeNumber(97))
	fmt.Println(IsPrimeNumber(1))
}
