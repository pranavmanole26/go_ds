package kthmissingnumber

func Find(nums []int, kth int) int {

	li := 0
	ri := len(nums) - 1

	if kth < nums[li] || kth > nums[ri] {
		return kth
	}

	for {
		if li > ri {
			break
		}

		mi := (li + ri) / 2
		//		4,	6,7,	   11
		//1,2,3,4,5,6,7,8,9,10,11
		//0,1,2,3,4,5,6,7,8,9,10
		//Here we need to check the difference between actual value at mi and the value at mi when no elements would be missing.
		//If that difference is less than the kth value then the kth value is at left side in the sequence.
		if nums[mi]-(mi+1) < kth {
			li = mi + 1
		} else {
			ri = mi - 1
		}
	}

	return kth + (ri + 1)
}
