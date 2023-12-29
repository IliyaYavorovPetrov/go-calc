package math

import "testing"

func TestSum(t *testing.T) {
	testCases := []struct {
		numbers  []int
		expected int
	}{
		{[]int{2, 3}, 5},
		{[]int{1, 2, 3, 4, 5}, 15},
	}

	for _, tc := range testCases {
		result := Sum(tc.numbers)
		if result != tc.expected {
			t.Errorf("Sum(%v) returned %d, expected %d", tc.numbers, result, tc.expected)
		}
	}
}

func TestSub(t *testing.T) {
	testCases := []struct {
		numbers  []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{10, 2}, 8},
		{[]int{100, 1, 2, 3, 4}, 90},
	}

	for _, tc := range testCases {
		result := Sub(tc.numbers)
		if result != tc.expected {
			t.Errorf("Sub(%v) returned %d, expected %d", tc.numbers, result, tc.expected)
		}
	}
}