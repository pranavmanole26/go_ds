package main

import "fmt"

func main() {
	// nums := []int{1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6}
	// nums := []int{1}
	// nums := []int{1, 2, 2, 3, 3}
	// nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6}
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	fmt.Println(getSingleElement(nums))
}

func getSingleElement(nums []int) (int, int) {
	if len(nums) == 1 {
		return nums[0], 0
	}

	if nums[0] != nums[1] {
		return nums[0], 0
	}

	lenN := len(nums)
	if nums[lenN-2] != nums[lenN-1] {
		return nums[lenN-1], lenN - 1
	}

	li := 1
	ri := lenN - 2

	for {
		if li > ri {
			return -1, -1
		}

		mi := (li + ri) / 2

		if nums[mi] != nums[mi+1] && nums[mi] != nums[mi-1] {
			return nums[mi], mi
		}

		// If mi is even then check values at index mi and mi+1
		// If mi is odd then check values at index mi and mi-1
		if mi%2 == 0 && nums[mi] == nums[mi+1] ||
			mi%2 == 1 && nums[mi] == nums[mi-1] {
			// Left half contains all double elements, hence eliminate left half
			li = mi + 1
		} else {
			// Right half contains all double elements, hence eliminate right half
			ri = mi - 1
		}
	}
}

// func searchSingle(sl []int) int {

// 	if len(sl) == 0 {
// 		return -1
// 	}

// 	if len(sl) == 1 {
// 		return sl[0]
// 	}

// 	if len(sl) == 2 {
// 		if sl[0] != sl[1] {
// 			return sl[0]
// 		} else {
// 			return -1
// 		}
// 	}

// 	li := 0
// 	ri := len(sl) - 1

// 	for {
// 		if li > ri {
// 			return -1
// 		}

// 		mi := (li + ri) / 2

// 		if mi == 0 && sl[mi] != sl[mi+1] ||
// 			mi == len(sl)-1 && sl[mi] != sl[mi-1] ||
// 			(sl[mi] != sl[mi-1] && sl[mi] != sl[mi+1]) {
// 			return sl[mi]
// 		}

// 		if mi%2 == 0 && sl[mi] == sl[mi+1] || mi%2 == 1 && sl[mi] == sl[mi-1] {
// 			li = mi + 1
// 		} else {
// 			ri = mi - 1
// 		}
// 	}
// }
