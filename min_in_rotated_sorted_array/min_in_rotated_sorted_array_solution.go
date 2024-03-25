package mininrotatedsortedarray

func Min(sl []int) int {
	li := 0
	ri := len(sl) - 1
	min := sl[0]

	for {
		if li > ri {
			return min
		}

		mi := (li + ri) / 2

		if sl[li] <= sl[mi] {
			if sl[li] < min {
				min = sl[li]
			}
			li = mi + 1
		} else {
			if sl[mi] < min {
				min = sl[mi]
			}
			ri = mi - 1
		}
	}
}
