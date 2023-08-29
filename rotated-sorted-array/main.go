package main

import (
	"fmt"
	"math"
)

func main() {
	// nums := []int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	// fmt.Println(findElement(nums, 10))

	// nums1 := []int{3, 1, 2, 3, 3}
	// nums1 := []int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	// fmt.Println(hasElementFromDuplicateElements(nums1, 10))

	nums2 := []int{9, 4}
	fmt.Println(minFromNonDuplicates(nums2))
}

// For non duplicate elements
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

			// Check if num is present in the right sorted array
			if nums[lI] <= num && num <= nums[mI] {
				// If yes then shift rI
				rI = mI - 1
			} else {
				// If no then shift lI
				lI = mI + 1
			}
		}
	}
}

func hasElementFromDuplicateElements(nums []int, num int) bool {
	lI := 0
	rI := len(nums) - 1

	for {
		if lI > rI {
			return false
		}
		mI := (lI + rI) / 2
		if nums[mI] == num {
			return true
		}
		// Check for this corner case where left, right and num could be the same values
		// If this is the case then trim down the array
		if nums[lI] == num && nums[rI] == num {
			lI = lI + 1
			rI = rI - 1
			continue
		}

		if nums[mI] < nums[rI] {
			if nums[mI] <= num && num <= nums[rI] {
				lI = mI + 1
			} else {
				rI = mI - 1
			}
		} else {
			if nums[lI] <= num && num <= nums[mI] {
				rI = mI - 1
			} else {
				lI = mI + 1
			}
		}
	}
}

func minFromNonDuplicates(nums []int) int {
	lI := 0
	rI := len(nums) - 1

	min := math.MaxInt

	for {
		if lI > rI {
			return min
		}
		mI := (lI + rI) / 2

		// First check sorted array
		if nums[mI] < nums[rI] {
			// Left part is sorted
			if min > nums[mI] {
				// If min is greater than the mI index value
				// then assign that value to min
				min = nums[mI]
			}
			// Eleminate left half
			rI = mI - 1
		} else {
			// Right part is sorted
			if min > nums[lI] {
				// If min is greater than the lI index value
				// then assign that value to min
				min = nums[lI]
			}
			// Elemenate right half
			lI = mI + 1
		}
	}
}
