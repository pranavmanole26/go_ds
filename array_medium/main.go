package main

import "fmt"

func UnionOfTwoArrays(first, second []int) []int {
	final := []int{}
	firstI := 0
	lastFirst := -1
	lastSecond := -1
	secondI := 0
	for {
		if firstI >= len(first) || secondI >= len(second) {
			break
		}
		if first[firstI] == lastFirst {
			firstI++
			continue
		}
		if second[secondI] == lastSecond {
			secondI++
			continue
		}
		if first[firstI] == lastSecond {
			firstI++
			continue
		}
		if second[secondI] == lastFirst {
			secondI++
			continue
		}
		if first[firstI] < second[secondI] {
			final = append(final, first[firstI], second[secondI])
			lastFirst = first[firstI]
			lastSecond = second[secondI]
			firstI++
			secondI++
		} else if second[secondI] < first[firstI] {
			final = append(final, second[secondI], first[firstI])
			lastFirst = first[firstI]
			lastSecond = second[secondI]
			firstI++
			secondI++
		} else {
			final = append(final, first[firstI])
			firstI++
			secondI++
		}

	}
	if firstI <= len(first) {
		for i := firstI; i < len(first); i++ {
			if first[i] != lastFirst {
				final = append(final, first[i])
			}
			lastFirst = first[i]
		}
	}
	if secondI <= len(second) {
		for i := secondI; i < len(second); i++ {
			if second[i] != lastSecond {
				final = append(final, second[i])
			}
			lastSecond = second[i]
		}
	}
	return final
}

func main() {
	fmt.Println(UnionOfTwoArrays([]int{1, 2, 2, 4, 6}, []int{2, 3, 4, 4, 4, 5, 7, 8}))
}
