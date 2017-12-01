package gostructures

import (
	"reflect"
	"testing"
)

func MaximumSubarray(array []int) (int, []int) {
	max := 0
	best := []int{}
	sum := 0
	current := []int{}
	for _, element := range array {
		sum += element
		current = append(current, element)
		if sum > max {
			max = sum
			best = current
		} else if sum < 0 {
			sum = 0
			current = []int{}
		}
	}
	return max, best
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
			actualSum, actualSubarray := MaximumSubarray(tt.array)
			if actualSum != tt.max {
				t.Errorf("max wrong: %v %v", actualSum, tt.max)
			}

			if !reflect.DeepEqual(actualSubarray, tt.subarray) {
				t.Errorf("subarray wrong: %v %v", actualSubarray, tt.subarray)
			}
		})
	}
}
