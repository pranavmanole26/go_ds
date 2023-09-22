package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(getPeakElement(nums))
}

func getPeakElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] > nums[1] {
		return nums[0]
	}
	n := len(nums)
	if nums[n-1] > nums[n-2] {
		return nums[n-1]
	}

	li := 1
	ri := n - 2

	for {
		// if li and ri crosses each other then return either of them
		if li > ri {
			return nums[li]
		}
		// find mid index
		mi := (li + ri) / 2
		// check if mid index element is peak
		if nums[mi] > nums[mi-1] && nums[mi] > nums[mi+1] {
			return nums[mi]
		} else if nums[mi] >= nums[mi-1] {
			// check if mid index element is greater than previous one
			// shift li, i.e. the numbers are growing towards right and the peak is at right
			li = mi + 1
		} else if nums[mi] >= nums[mi+1] {
			// shift ri, i.e. the numbers are growing towards left and the peak is at left
			ri = mi - 1
		}
		// else {
		// 	// In case of multiple peaks we can shift to either side
		// 	li = mi + 1
		// }
	}
}
