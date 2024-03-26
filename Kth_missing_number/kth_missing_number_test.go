package kthmissingnumber

import "testing"

func TestFind(t *testing.T) {
	testcases := []struct {
		name           string
		sl             []int
		kth            int
		expectedOutput int
	}{
		{
			name:           "Test 1",
			sl:             []int{2, 3, 4, 7, 11},
			kth:            5,
			expectedOutput: 9,
		},
		{
			name:           "Test 2",
			sl:             []int{4, 7, 11},
			kth:            3,
			expectedOutput: 3,
		},
		{
			name:           "Test 3",
			sl:             []int{4, 7, 11},
			kth:            12,
			expectedOutput: 12,
		},
		{
			name:           "Test 4",
			sl:             []int{4, 6, 7, 11},
			kth:            6,
			expectedOutput: 9,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			output := Find(tc.sl, tc.kth)
			if output != tc.expectedOutput {
				t.Errorf("Expected=%d, Got=%d\n", tc.expectedOutput, output)
			}
		})
	}
}
