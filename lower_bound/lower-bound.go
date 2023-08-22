package main

func lowerBound(nums []int, num int) int {
	lI := 0
	rI := len(nums) - 1

	lowerBnd := len(nums)

	for {
		if lI > rI {
			return lowerBnd
		}
		mI := (lI + rI) / 2
		if nums[mI] >= num {
			lowerBnd = mI
		}
		if nums[mI] > num {
			rI = mI - 1
		} else {
			lI = mI + 1
		}
	}
}
