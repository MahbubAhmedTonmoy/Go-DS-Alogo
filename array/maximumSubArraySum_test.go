package array

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		expected int
	}{
		{
			name:     "Empty slice",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "Max is 0",
			slice:    []int{0, -1, -2, -4, -5},
			expected: 0,
		},
		{
			name:     "Max is -1",
			slice:    []int{-1, -3, -2, -5, -7},
			expected: -1,
		},
		{
			name:     "Max is 7",
			slice:    []int{-2, -5, 6, 0, -2, 0, -3, 1, 0, 5, -6},
			expected: 7,
		},
		{
			name:     "Max is 11",
			slice:    []int{1, 7, -2, -5, 10, -1},
			expected: 11,
		},
		{
			name:     "Max is 11",
			slice:    []int{-2, 11, -4, 2, -3, -10},
			expected: 11,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			if result := MaximumSubArraySum(test.slice); result != test.expected {
				t.Fatalf("%s\n\tslice: %v, expected: %v, returned: %v", test.name, test.slice, test.expected, result)
			}
		})
	}
}
