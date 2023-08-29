package main

import "fmt"

func main() {
	nums := []int{7, 8, 9, 1, 2, 3, 4, 5, 6}

	fmt.Println(findElement(nums, 10))
}

func findElement(nums []int, num int) int {
	lI := 0
	rI := len(nums) - 1

	for {
		if lI > rI {
			return -1
		}
		mI := (lI + rI) / 2
		if nums[mI] == num {
			return mI
		}

		// Find sorted array
		if nums[mI] < nums[rI] {
			// Left array is sorted

			// Check if num is present in the left sorted array
			if nums[mI] <= num && num <= nums[rI] {
				// If num is in the left sorted array then shift lI
				lI = mI + 1
			} else {
				// If num is not in the left sorted array then shift rI
				rI = mI - 1
			}
		} else {
			// Right array is sorted

			if nums[lI] <= num && num <= nums[mI] {
				rI = mI - 1
			} else {
				lI = mI + 1
			}
		}
	}
}
