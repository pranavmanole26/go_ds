package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSquareRoot(t *testing.T) {

	testCases := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "TC 1",
			input:  16,
			output: 4,
		},
		{
			name:   "TC 2",
			input:  17,
			output: 4,
		},
		{
			name:   "TC 3",
			input:  24,
			output: 4,
		},
		{
			name:   "TC 4",
			input:  0,
			output: 1,
		},
	}

	assertion := assert.New(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := getSquareRoot(tc.input)
			assertion.Equal(tc.output, actualOutput)
		})
	}
}
