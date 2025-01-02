package main

import (
	"fmt"
	"math"
)

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

// Question: https://takeuforward.org/data-structure/stock-buy-and-sell/
func BuyAndSellStocks(nums []int) int {
	maxProfit := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] && nums[j]-nums[i] > maxProfit {
				maxProfit = nums[j] - nums[i]
			}
		}
	}
	return maxProfit
}

// **** Optimal solution
func BuyAndSellStocksOptimal(nums []int) int {
	maxProfit := 0
	minPrice := math.MaxInt
	for i := 0; i < len(nums); i++ {
		// Update minPrice value everytime we get lesser value
		// And go for next iteration. No need to perform any further action in this iteration
		if minPrice > nums[i] {
			minPrice = nums[i]
			continue
		}
		// If maxProfit is lesser than the difference between current value and minPrice then update maxProfit
		// I.e. update maxProfit if the current profit is higher.
		//  The current profit is the difference between current value and the minPrice
		if maxProfit < nums[i]-minPrice {
			maxProfit = nums[i] - minPrice
		}
	}
	return maxProfit
}

func isPositive(num int) bool {
	return num > 0
}

func isNegative(num int) bool {
	return num < 0
}

func isOppositeSigns(num1, nums2 int) bool {
	return (isNegative(num1) && isPositive(nums2)) || (isPositive(num1) && isNegative(nums2))
}

// Question: https://takeuforward.org/arrays/rearrange-array-elements-by-sign/
func RearrageBySigns(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// The sortedI help us to identify till what index an array is sorted
	sortedI := 0
	for i := 1; i < len(nums)-1; i++ {
		if isOppositeSigns(nums[sortedI], nums[i]) {
			// Do swapping only when the current index and the sortedI are having difference more than one
			// Meaning current index is far away from the sortedI because of all the consecutive numbers are of same sign
			if i-sortedI > 1 {
				nums[sortedI+1], nums[i] = nums[i], nums[sortedI+1]
				// After swapping if the number at sorted+1 and number next to sorted+1 are of different sign then increment sortedI by 2
				if isOppositeSigns(nums[sortedI+1], nums[sortedI+2]) {
					sortedI++
				}
			}
			sortedI++
		}
	}
	return nums
}

// Question: https://takeuforward.org/data-structure/longest-consecutive-sequence-in-an-array/
// Works well with +ve numbers
// TODO: need to use set to have unique numbers
func LongestConsecutiveSequence(nums []int) int {
	longestConsecutiveSubArray := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == -100 {
			continue
		}
		curSubArray := []int{nums[i]}
		for j := i + 1; j < len(nums); {
			if nums[i] == nums[j] {
				nums[j] = -100
				j++
				continue
			}
			// fmt.Println(curSubArray)
			if curSubArray[len(curSubArray)-1]+1 == nums[j] {
				curSubArray = append(curSubArray, nums[j])
				nums[j] = -100
				j = i + 1
			} else if curSubArray[len(curSubArray)-1]-1 == nums[j] ||
				curSubArray[0]-1 == nums[j] {
				curSubArray = append([]int{nums[j]}, curSubArray...)
				nums[j] = -100
				j = i + 1
			} else {
				j++
			}
		}
		if len(curSubArray) > longestConsecutiveSubArray {
			longestConsecutiveSubArray = len(curSubArray)
		}
	}
	return longestConsecutiveSubArray
}

// func NextLexicographicalSequence(nums []int) []int {
// 	mp := map[int]struct{}{}
// 	for i := 0; i < len(nums); i++ {
// 		_, ok1 := mp[nums[i]-1]
// 		_, ok2 := mp[nums[i]+1]
// 		if ok1 || ok2 {

// 		}
// 	}
// }

func SetMatrixZero(matrix [][]int) [][]int {
	zeroI := map[int]struct{}{}
	zeroJ := map[int]struct{}{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeroI[i] = struct{}{}
				zeroJ[j] = struct{}{}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_, isIZero := zeroI[i]
			_, isJZero := zeroJ[j]
			if isIZero || isJZero {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func RotateMatrixByNintyDegree(matrix [][]int) [][]int {
	rotatedMatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		row := make([]int, len(matrix[i]))
		rotatedMatrix[i] = row
	}
	col := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			rotatedMatrix[j][col] = matrix[i][j]
		}
		col--
	}
	return rotatedMatrix
}

// 	0 1 2   0 1 2
// 0	1 2 3   7 4 1
// 1	4 5 6	8 5 2
// 2	7 8 9	9 6 3

// 1: [0,0] -> [0,2]
// 2: [0,1] -> [1,2]
// 3: [0,2] -> [2,2]

// 4: [1,0] -> [0,1]
// 5: [1,1] -> [1,1]
// 6: [1,2] -> [2,1]

// 7:
// 8:
// 9:

func main() {
	// fmt.Println(UnionOfTwoArrays([]int{1, 2, 2, 4, 6}, []int{2, 3, 4, 4, 4, 5, 7, 8}))
	// fmt.Println(FindSingleElement([]int{1, 1, 2, 3, 4, 3, 2}))
	// fmt.Println(LongestSubArrayWithSum([]int{2, 3, 5, 1, 9}, 10))
	// fmt.Println(LongestSubArrayWithSum([]int{-1, 1, 1}, 1))
	// fmt.Println(TwoSum([]int{2, 6, 5, 8, 11}, 14))
	// fmt.Println(SortZeroOneTwo([]int{2, 0, 2, 1, 2, 1, 0, 1}))
	// fmt.Println(BuyAndSellStocks([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(BuyAndSellStocksOptimal([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(BuyAndSellStocksOptimal([]int{7, 5, 5, 3, 2, 1}))
	// fmt.Println(BuyAndSellStocksOptimal([]int{7, 3, 5, 1, 6, 4}))
	// fmt.Println(RearrageBySigns([]int{1, 2, -4, -5}))
	// fmt.Println(RearrageBySigns([]int{1, 2, -3, -1, 2, -3}))
	// fmt.Println(RearrageBySigns([]int{1, 2, 3, -1, -2, -3}))
	// fmt.Println(RearrageBySigns([]int{1, 2, 3, 4, -4, -1, -2, -3}))
	// fmt.Println(LongestConsecutiveSequence([]int{100, 200, 1, 3, 2, 4}))
	// fmt.Println(LongestConsecutiveSequence([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	// fmt.Println(LongestConsecutiveSequence([]int{100, 4, 200, 1, 3, 2}))
	// fmt.Println(LongestConsecutiveSequence([]int{1, 2, 0, 1}))
	// fmt.Println(SetMatrixZero([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	// fmt.Println(SetMatrixZero([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}))
	fmt.Println(RotateMatrixByNintyDegree([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(RotateMatrixByNintyDegree([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}))
}
