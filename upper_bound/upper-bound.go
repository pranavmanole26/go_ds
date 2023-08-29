package main

func upperBound(nums []int, num int) int {
	lI := 0
	rI := len(nums) - 1
	upperBnd := len(nums)

	for {
		if lI > rI {
			return upperBnd
		}
		cI := (lI + rI) / 2

		if nums[cI] > num {
			upperBnd = cI
			rI = cI - 1
		} else {
			lI = cI + 1
		}
	}
}
