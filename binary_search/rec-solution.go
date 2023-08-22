package main

func search(nums []int, lI int, rI int, num int) int {
	if lI > rI {
		return -1
	}
	cI := (lI + rI) / 2
	if nums[cI] == num {
		return cI
	}
	if nums[cI] > num {
		rI = cI - 1
	} else {
		lI = cI + 1
	}
	return search(nums, lI, rI, num)
}
