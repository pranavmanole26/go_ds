package mincapacitytoloadship

import "testing"

func TestFind(t *testing.T) {
	testcases := []struct {
		name        string
		loads       []int
		maxDays     int
		expectedCap int
	}{
		{
			name:        "Test 1",
			loads:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			maxDays:     5,
			expectedCap: 15,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			cap := Find(tc.loads, tc.maxDays)

			if cap != tc.expectedCap {
				t.Errorf("Expected=%d, Got=%d\n", tc.expectedCap, cap)
			}
		})
	}
}
