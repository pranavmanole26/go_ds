package binarysearch

func BinSearchRec(nums []int, lI int, rI int, num int) int {
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
	return BinSearchRec(nums, lI, rI, num)
}

func BinSearch(sl []int, target int) int {
	l := 0
	r := len(sl) - 1

	for {
		if l > r {
			return -1
		}
		m := (l + r) / 2

		if sl[m] == target {
			return m
		} else if sl[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
}
