package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPeakElement(t *testing.T) {
	assertion := assert.New(t)

	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "First element peak",
			input:  []int{3, 2, 1},
			output: 3,
		},
		{
			name:   "Last element peak",
			input:  []int{1, 2, 3},
			output: 3,
		},
		{
			name:   "Single element",
			input:  []int{1},
			output: 1,
		},
		{
			name:   "Two elements 1",
			input:  []int{2, 3},
			output: 3,
		},
		{
			name:   "Two elements 2",
			input:  []int{3, 2},
			output: 3,
		},
		{
			name:   "Inside peak",
			input:  []int{1, 2, 3, 4, 3, 2},
			output: 4,
		},
		{
			name:   "All peak",
			input:  []int{3, 3, 3},
			output: 3,
		},
		// {
		// 	name:   "Duplicate consequitive peaks",
		// 	input:  []int{1, 2, 3, 3, 2},
		// 	output: 3,
		// },
		// {
		// 	name:   "Multiple peaks",
		// 	input:  []int{1, 5, 1, 2, 1},
		// 	output: 2,
		// },
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			actualOutput := getPeakElement(tC.input)
			assertion.Equal(tC.output, actualOutput)
		})
	}

}

func BenchmarkGetPeakElement(b *testing.B) {
	nums := []int{1, 2, 3, 4, 3, 2}
	for i := 0; i < b.N; i++ {
		getPeakElement(nums)
	}
}
