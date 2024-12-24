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

func FindMaxConsecutiveOfGivenNumber(nums []int, num int) int {
	maxCount := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			count++
			if maxCount < count {
				maxCount = count
			}
		} else {
			count = 0
		}
	}
	return maxCount
}

// Question: https://takeuforward.org/data-structure/find-the-majority-element-that-occurs-more-than-n-2-times/
func MajorityElement(nums []int) int {
	occuranceMap := map[int]int{}
	for _, n := range nums {
		occuranceMap[n] = occuranceMap[n] + 1
	}
	for k, v := range occuranceMap {
		if v >= len(nums)/2 {
			return k
		}
	}
	return -1
}

// Not working with fmt.Println(MajorityElementOptimized([]int{1, 1, 2, 2, 1, 2, 3, 3, 2, 2, 1}))
// Giving wrong answer as -1
// func MajorityElementOptimized(nums []int) int {
// 	var el int
// 	var count int
// 	for _, num := range nums {
// 		if count == 0 {
// 			count = 1
// 			el = num
// 		} else if el == num {
// 			count++
// 		} else {
// 			count--
// 		}
// 	}
// 	count = 0
// 	for _, n := range nums {
// 		if n == el {
// 			count++
// 		}
// 	}
// 	if count >= len(nums)/2 {
// 		return el
// 	}

// 	return -1
// }

func MaximumSubarraySum(nums []int) int {
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
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
	// fmt.Println(FindTheMissingNumberFromSeries([]int{1, 2, 3, 5, 6}))
	// fmt.Println(FindMaxConsecutiveOfGivenNumber([]int{1, 1, 2, 2, 1, 1, 1, 2, 3, 1, 1, 1, 1}, 2))
	// fmt.Println(MajorityElement([]int{1, 1, 2, 2, 1, 2, 3, 3, 2, 2, 1}))
	fmt.Println(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

}
