package main

import "fmt"

func main() {
	fmt.Println(getPageFaults([]int{7, 0, 1, 2, 0, 3, 0, 4, 2, 3, 0, 3, 2, 3}, 4))
}

func getPageFaults(pages []int, capacity int) int {
	set := make(map[int]struct{})
	uses := make(map[int]int)
	pageFault := 0

	for i := 0; i < len(pages); i++ {
		if len(set) < capacity {
			// If set is not full
			if _, ok := set[pages[i]]; !ok {
				set[pages[i]] = struct{}{}
				pageFault++
			}
			uses[pages[i]] = i
		} else {
			// If set is full
			if _, ok := set[pages[i]]; !ok {
				set[pages[i]] = struct{}{}
				pageFault++

				lruI, lru := i-0, uses[i-0]

				for k, v := range uses {
					if v < lruI {
						lruI = v
						lru = k
					}
				}

				delete(set, lru)
				delete(uses, lru)
				set[pages[i]] = struct{}{}
			}
			uses[pages[i]] = i

		}
	}

	return pageFault
}
