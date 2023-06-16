package main

import "fmt"

func main() {
	fmt.Println(getPageFaults([]int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2}, 4)) // 6
	fmt.Println(getPageFaults([]int{}, 4))                                      // 0
	fmt.Println(getPageFaults([]int{7}, 4))                                     // 1
	fmt.Println(getPageFaults([]int{1, 1, 1, 1, 1, 1, 1}, 4))                   // 1
	fmt.Println(getPageFaults([]int{1, 2, 3, 4, 5, 1, 1}, 4))                   // 6
}

func getPageFaults(pages []int, capacity int) int {
	set := make(map[int]struct{ index int })
	pageFault := 0

	for i := 0; i < len(pages); i++ {
		if len(set) < capacity {
			// If set is not full
			if _, ok := set[pages[i]]; !ok {
				pageFault++
			}
		} else {
			// If set is full
			if _, ok := set[pages[i]]; !ok {
				pageFault++

				// LRU algo start
				lruI, lru := i-0, 0

				for k, v := range set {
					if v.index < lruI {
						lruI = v.index
						lru = k
					}
				}
				delete(set, lru)
				// LRU algo End
			}
		}
		set[pages[i]] = struct{ index int }{i}
	}

	return pageFault
}
