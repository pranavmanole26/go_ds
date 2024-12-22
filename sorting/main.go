package main

import "fmt"

// In the Selection Sort algorithm we use a nested loop
// 1. The outer loop helps to decide range of a given array. The range gets reduced by one index in each iteration.
// Meaning if first iteration loops over from 0 to last index then the second will iterate from 1 to last and so on.
// 2. The inner loop helps to actually iterate over the subset obtained in the each outer iteration.
// In each inner loop iteration we check what is the smallest element and updates the "minI" variable with the index of it.
// 3. After inner loop iteration finishes we swaps the first element in the subset with the smallest element.
func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		// min := nums[i]
		minI := i
		for j := i; j < len(nums); j++ {
			if nums[minI] > nums[j] {
				// min = nums[j]
				minI = j
			}
		}
		nums[i], nums[minI] = nums[minI], nums[i]
	}
	return nums
}

// In the Bubble Sort nested loop is used
// 1. The outer loop is usefull to decide the range of an array. In each iteration range is reduced by one index.
// The subset always starts from the first index and ends with the index lesser than the preious last index.
// I.e. In the first iteration first index is 0 and the last index is "len(nums) - 1".
// In the second iteration the first index is 0 and the last index becomes "len(nums) - 2"
// 2. The inner loop iterates over the subset of an array obtained in the each outer loop iteration.
// In the each iteration we check which element is bigger between adjucent elements. The bigger element is pushed to the right.
// In this way in each inner loop iteration we push bigger elements to the right.
func BubbleSort(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// In the Insertion Sort we need to take help of nested for loops
// Outer for loop will decide for the range of a sub-array. In a first outer loop iteration sub-array is from index 0 to index 1
// In the inner loop iteration starting from the last element from the sub-array comparison happens between the previous elements
// And if the element is smaller then it is moved to rightside in the sub-array.
// In this way the sub-arrays are sorted in each iteration of the inner for loop.
func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				fmt.Println("iteration: ", i, ".", j, ":", nums)
			}
		}
	}
	return nums
}

func main() {
	nums := []int{4, 3, 5, 1, 2}
	// fmt.Println(SelectionSort(nums))
	// fmt.Println(BubbleSort(nums))
	fmt.Println(InsertionSort(nums))
}
