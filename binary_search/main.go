package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	i := search(nums, 0, len(nums)-1, 1)
	fmt.Println(i)
}
