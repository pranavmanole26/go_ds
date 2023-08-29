package main

import "fmt"

func main() {
	nums := []int{0, 2, 4, 6, 6, 8, 10}
	num := 11
	fO := firstOccurrence(nums, num)
	// First check first occurrence. If first occurrence is not present then last occurrence won't be present.
	// Hence, return -1, -1.
	if fO == -1 {
		fmt.Println(-1, -1)
		return
	}
	lO := lastOccurrence(nums, num)
	fmt.Println(fO, lO)
}

func firstOccurrence(nums []int, num int) int {
	first := -1

	lI := 0
	rI := len(nums) - 1

	for {
		if lI > rI {
			return first
		}

		mI := (lI + rI) / 2

		if nums[mI] == num {
			first = mI
			rI = mI - 1
		} else if nums[mI] < num {
			lI = mI + 1
		} else {
			rI = mI - 1
		}
	}
}

func lastOccurrence(nums []int, num int) int {
	lI := 0
	rI := len(nums) - 1
	last := -1

	for {
		if lI > rI {
			return last
		}
		mI := (lI + rI) / 2
		if nums[mI] == num {
			last = mI
			lI = mI + 1
		} else if nums[mI] > num {
			rI = mI - 1
		} else {
			lI = mI + 1
		}
	}
}
