package main

import "fmt"

// Question: https://takeuforward.org/data-structure/print-all-divisors-of-a-given-number/

// Optimal solution
// Considering divisor and n/divisor both are divisors to the n. Hence appending both to the slice in the same iteration.
// When divisor becomes greater than or equal to the n/divisor then breaking the loop and optimising the time complexity.
func GetAllDivisors(n int) []int {
	allDivisors := []int{1, n}
	divisor := 2
	for {
		if n%divisor == 0 {
			if divisor >= n/divisor {
				if divisor == n/divisor {
					allDivisors = append(allDivisors, divisor)
				}
				break
			}
			allDivisors = append(allDivisors, divisor, n/divisor)
		}
		divisor++
	}

	return allDivisors
}

func main() {
	fmt.Println(GetAllDivisors(20))
	fmt.Println(GetAllDivisors(36))
	fmt.Println(GetAllDivisors(45))
	fmt.Println(GetAllDivisors(100))
	// fmt.Println(GetAllDivisors(1000000000))
}
