package binarysearch

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
