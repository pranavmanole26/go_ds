package main

import (
	"fmt"
	"math"
)

func main() {
	// nums := []int{4, 5, 1, 2, 3}
	// nums := []int{1, 2, 3, 4, 5}
	// nums := []int{5, 4, 3, 2, 1}
	nums := []int{7, 5, 4, 2, 3}

	rotations := numberOfRotations(nums)
	fmt.Println(rotations)
}

func numberOfRotations(nums []int) int {
	return getMinValueIndex(nums)
}

func getMinValueIndex(nums []int) int {

	min := math.MaxInt
	minI := -1

	li := 0
	ri := len(nums) - 1

	for {
		if li > ri {
			return minI
		}

		mi := (li + ri) / 2
		// Find sorted half
		if nums[mi] < nums[ri] {
			// right side is sorted
			if nums[mi] < min {
				min = nums[mi]
				minI = mi
			}
			ri = mi - 1
		} else {
			// left side is sorted
			if nums[li] < min {
				min = nums[li]
				minI = li
			}
			li = mi + 1
		}
	}
}
