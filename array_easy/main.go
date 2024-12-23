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

func ShiftArrayElementsByKPlacesToLeft(nums []int, k int) []int {
	temp := make([]int, k)
	for i := 0; i < k; i++ {
		temp[i] = nums[i]
	}
	temp1 := make([]int, len(nums)-k, len(nums))
	for i := 0; i < len(nums)-k; i++ {
		temp1[i] = nums[i+k]
	}
	temp1 = append(temp1, temp...)
	return temp1
}

func ShiftAllOccurancesOfGivenNumberToTheLeft(nums []int, num int) []int {
	numI := 0
	firstNum := false
	for i := 0; i < len(nums); i++ {
		// This if must execute first in the iteration in order to decide the first occurance of "num" and store it's index in "numI"
		// Once we get the index we should go for next iteration
		if nums[i] == num && !firstNum {
			numI = i
			firstNum = true
			continue
		}
		// Once we get the first occurance of the "num" we need to check if the next element is not the "num"
		// If it is not "num" then swap these two numbers and increment numI by one
		if nums[i] != num && firstNum {
			nums[i], nums[numI] = nums[numI], nums[i]
			numI++
		}
	}
	return nums
}

func FindTheMissingNumberFromSeries(nums []int) int {
	actualSum := 0
	// Find actual sum
	for _, num := range nums {
		actualSum += num
	}
	// Find sum of ideal series
	// Sn = n/2 * (2*a + (n-1)*d)
	sumOfIdealSeries := (len(nums) + 1) / 2 * (2*nums[0] + (nums[len(nums)-1]-1)*1)
	// The difference between ideal sum and actual sum is the missing number
	return sumOfIdealSeries - actualSum
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
	// fmt.Println(RemoveDuplicatesFromSorted([]int{1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 6}))
	// fmt.Println(ShiftArrayElementsByKPlacesToLeft([]int{1, 2, 3, 4, 5}, 2))
	// fmt.Println(ShiftAllOccurancesOfGivenNumberToTheLeft([]int{1, 2, 3, 4, 2, 3, 2, 1, 2}, 3))
	fmt.Println(FindTheMissingNumberFromSeries([]int{1, 2, 3, 5, 6}))
}
