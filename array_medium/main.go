package main

import "fmt"

func UnionOfTwoArrays(first, second []int) []int {
	final := []int{}
	firstI := 0
	lastFirst := -1
	lastSecond := -1
	secondI := 0
	for {
		if firstI >= len(first) || secondI >= len(second) {
			break
		}
		if first[firstI] == lastFirst {
			firstI++
			continue
		}
		if second[secondI] == lastSecond {
			secondI++
			continue
		}
		if first[firstI] == lastSecond {
			firstI++
			continue
		}
		if second[secondI] == lastFirst {
			secondI++
			continue
		}
		if first[firstI] < second[secondI] {
			final = append(final, first[firstI], second[secondI])
			lastFirst = first[firstI]
			lastSecond = second[secondI]
			firstI++
			secondI++
		} else if second[secondI] < first[firstI] {
			final = append(final, second[secondI], first[firstI])
			lastFirst = first[firstI]
			lastSecond = second[secondI]
			firstI++
			secondI++
		} else {
			final = append(final, first[firstI])
			firstI++
			secondI++
		}

	}
	if firstI <= len(first) {
		for i := firstI; i < len(first); i++ {
			if first[i] != lastFirst {
				final = append(final, first[i])
			}
			lastFirst = first[i]
		}
	}
	if secondI <= len(second) {
		for i := secondI; i < len(second); i++ {
			if second[i] != lastSecond {
				final = append(final, second[i])
			}
			lastSecond = second[i]
		}
	}
	return final
}

// Question: https://takeuforward.org/arrays/find-the-number-that-appears-once-and-the-other-numbers-twice/
func FindSingleElement(nums []int) int {
	mapOfOccurance := map[int]struct{}{}
	for _, n := range nums {
		if _, ok := mapOfOccurance[n]; !ok {
			mapOfOccurance[n] = struct{}{}
		} else {
			delete(mapOfOccurance, n)
		}
	}
	for k := range mapOfOccurance {
		return k
	}
	return -1
}

// Question: https://takeuforward.org/data-structure/longest-subarray-with-given-sum-k/
// TODO: Understand the hashing solution
func LongestSubArrayWithSum(nums []int, sum int) []int {
	longestSubArray := []int{}
	for i := 0; i < len(nums); i++ {
		curSum := 0
		subArray := []int{}
		for j := i; j < len(nums); j++ {
			curSum += nums[j]
			subArray = append(subArray, nums[j])
			if curSum == sum && len(longestSubArray) < len(subArray) {
				longestSubArray = subArray
				break
			}
			if curSum > sum {
				break
			}
		}
	}
	return longestSubArray
}

func TwoSum(nums []int, sum int) (int, int) {
	mapRemainingToNumber := map[int]int{}

	for _, num := range nums {
		mapRemainingToNumber[sum-num] = num
	}
	for k, v := range mapRemainingToNumber {
		if val, ok := mapRemainingToNumber[sum-k]; ok {
			return v, val
		}
	}
	return -1, -1
}

// **** Understand this solution again and again ****
// Question: https://takeuforward.org/data-structure/sort-an-array-of-0s-1s-and-2s/
func SortZeroOneTwo(nums []int) []int {
	left := 0
	mid := 0
	right := len(nums) - 1

	for {
		if mid >= right {
			break
		}
		if nums[mid] == 0 {
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[right], nums[mid] = nums[mid], nums[right]
			right--
		}
	}

	return nums
}

func main() {
	// fmt.Println(UnionOfTwoArrays([]int{1, 2, 2, 4, 6}, []int{2, 3, 4, 4, 4, 5, 7, 8}))
	// fmt.Println(FindSingleElement([]int{1, 1, 2, 3, 4, 3, 2}))
	fmt.Println(LongestSubArrayWithSum([]int{2, 3, 5, 1, 9}, 10))
	fmt.Println(LongestSubArrayWithSum([]int{-1, 1, 1}, 1))
	// fmt.Println(TwoSum([]int{2, 6, 5, 8, 11}, 14))
	// fmt.Println(SortZeroOneTwo([]int{2, 0, 2, 1, 2, 1, 0, 1}))
}
