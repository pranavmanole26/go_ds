package main

import (
	"fmt"
	binarysearch "go_ds/binary_search"
)

var (
	binTestData = []int{1, 2, 3, 4, 5}
)

func main() {
	i := binarysearch.BinSearch(binTestData, 3)

	fmt.Println(i)
}
