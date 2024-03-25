package mininrotatedsortedarray

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	testcases := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "Test1",
			input:          []int{6, 7, 8, 1, 2, 3, 4, 5},
			expectedOutput: 1,
		},
		{
			name:           "Test2",
			input:          []int{6, 8, 2, 4, 5},
			expectedOutput: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			op := Min(tc.input)
			if op != tc.expectedOutput {
				fmt.Println(fmt.Errorf("Expected -> %d. Got -> %d", tc.expectedOutput, op))
			}
		})
	}
}
