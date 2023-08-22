package main

import "fmt"

func main() {
	nums := []int{0, 2, 4, 6, 6, 8, 10}
	fmt.Println(lowerBound(nums, -1))
	fmt.Println(lowerBound(nums, 1))
	fmt.Println(lowerBound(nums, 3))
	fmt.Println(lowerBound(nums, 6))
	fmt.Println(lowerBound(nums, 10))
	fmt.Println(lowerBound(nums, 11))
}
