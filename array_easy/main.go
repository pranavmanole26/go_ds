package main

import "fmt"

func LargestElement(nums []int) int {
	largest := 0

	for _, n := range nums {
		if largest < n {
			largest = n
		}
	}
	return largest
}

func SecondSmallestAndSecondLargest(nums []int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}
	if len(nums) == 1 {
		return nums[0], nums[0]
	}
	smallest, secondSmallest := nums[0], nums[0]
	largest, secondLargest := nums[0], nums[0]

	for _, n := range nums {
		if n < smallest {
			// Update secondSmallest by earlier smallest
			secondSmallest = smallest
			smallest = n
		}
		if n < secondSmallest && n > smallest {
			secondSmallest = n
		}
		if n > largest {
			// Update secondLargest by earlier largest
			secondLargest = largest
			largest = n
		}
		if n > secondLargest && n < largest {
			secondLargest = n
		}
	}
	return secondSmallest, secondLargest
}

func IsSortedBF(nums []int) bool {
	// Below two if blocks are the only conditions where we can surely say the sub-array is sorted
	if len(nums) <= 1 {
		return true
	}
	if len(nums) == 2 {
		return nums[0] < nums[1]
	}
	mid := (len(nums) - 1) / 2
	if nums[0] > nums[mid] || nums[len(nums)-1] < nums[mid] {
		return false
	}
	isLeftSorted := IsSortedBF(nums[:mid])
	isRightSorted := IsSortedBF(nums[mid:])
	return isLeftSorted && isRightSorted
}

func RemoveDuplicatesFromSorted(nums []int) []int {
	for i := 0; i < len(nums)-1; {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i+1], nums[i+2:]...)
		} else {
			i++
		}
	}
	return nums
}

func main() {
	// fmt.Println(LargestElement([]int{2, 4, 5, 3, 1}))
	// fmt.Println(SecondSmallestAndSecondLargest([]int{2, 4, 5, 3, 1}))
	// fmt.Println(SecondSmallestAndSecondLargest([]int{4, 5, 3, 1, 6, 7, 8, 10, 9, 7, 2}))
	// fmt.Println(IsSortedBF([]int{1, 2, 3, 4, 5, 7, 6, 8, 9, 10, 11, 12, 14, 16, 20}))
	// fmt.Println(IsSortedBF([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 16, 20}))
	// 1,2,2,3
	// i 	-> 1
	// i+1 -> 2
	// nums[:i+1] -> 1,2
	// nums[i+1:] -> 3
	fmt.Println(RemoveDuplicatesFromSorted([]int{1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 6}))
}
