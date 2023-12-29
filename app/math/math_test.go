package math

import "testing"

func TestSum(t *testing.T) {
	testCases := []struct {
		numbers  []int
		expected int
	}{
		{[]int{2, 3}, 5},
		{[]int{1, 2, 3, 4, 5}, 14},
	}

	for _, tc := range testCases {
		result := Sum(tc.numbers)
		if result != tc.expected {
			t.Errorf("Add(%v) returned %d, expected %d", tc.numbers, result, tc.expected)
		}
	}
}