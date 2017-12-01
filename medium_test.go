package gostructures

import "testing"

func MaximumSubarray(array []int) int {
	max := 0
	sum := 0
	for _, element := range array {
		sum += element
		if sum > max {
			max = sum
		} else if sum < 0 {
			sum = 0
		}
	}
	return max
}

func TestMaximumSubarray(t *testing.T) {
	testcases := []struct {
		name     string
		array    []int
		max      int
		subarray []int
	}{
		{
			name:     "example",
			array:    []int{2, -8, 3, -2, 4, -10},
			max:      5,
			subarray: []int{3, -2, 4},
		},
		{
			name:     "first value only",
			array:    []int{100, -1000, 1},
			max:      100,
			subarray: []int{100},
		},

		{
			name:     "last value only",
			array:    []int{1, -1000, 100},
			max:      100,
			subarray: []int{100},
		},

		{
			name:     "middle value only",
			array:    []int{1, -1000, 100, -1000, 1},
			max:      100,
			subarray: []int{100},
		},

		{
			name:     "everything",
			array:    []int{10, -1, -2, -3, 10, -4, -5, 10},
			max:      15,
			subarray: []int{10, -1, -2, -3, 10, -4, -5, 10},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			actual := MaximumSubarray(tt.array)
			if actual != tt.max {
				t.Errorf("max wrong: %v %v", actual, tt.max)
			}
		})
	}
}
