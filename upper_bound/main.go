package main

import "fmt"

func main() {
	nums := []int{0, 2, 4, 6, 6, 8, 10}
	fmt.Println(upperBound(nums, -1))
	fmt.Println(upperBound(nums, 1))
	fmt.Println(upperBound(nums, 3))
	fmt.Println(upperBound(nums, 6))
	fmt.Println(upperBound(nums, 10))
	fmt.Println(upperBound(nums, 11))
}
