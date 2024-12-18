package main

import (
	"fmt"
	"math"
)

func HighestAndLowestFrequency(nums []int) (int, int) {
	mapNums := make(map[int]int)

	for _, n := range nums {
		mapNums[n] = mapNums[n] + 1
	}

	highestFreq, highestNum := 0, 0
	lowestFreq, lowestNum := int(math.Pow(10, 9)), 0

	for k, v := range mapNums {
		if v > highestFreq {
			highestFreq = v
			highestNum = k
		}
		if v < lowestFreq {
			lowestFreq = v
			lowestNum = k
		}
	}
	return highestNum, lowestNum
}

func main() {
	fmt.Println(HighestAndLowestFrequency([]int{1, 2, 1, 1, 3, 2}))
	fmt.Println(HighestAndLowestFrequency([]int{1, 2, 1, 1, 3, 4, 3, 2}))
	fmt.Println(HighestAndLowestFrequency([]int{1, 2, 1, 1, 3, 4, 3, 2}))
}
