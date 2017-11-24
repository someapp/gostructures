package gostructures

import (
	"reflect"
	"testing"
)

type sortingTestCase struct {
	name            string
	input, expected []int
}

func sortingTestCases() []sortingTestCase {
	return []sortingTestCase{
		{
			name:     "in order",
			input:    []int{0, 1, 2, 3, 4, 5, 6, 7},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "reversed",
			input:    []int{7, 6, 5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "shuffled",
			input:    []int{5, 0, 2, 4, 7, 3, 6, 1},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range sortingTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			actual := BubbleSort(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v to give: %v, actual: %v", tt.input, tt.expected, actual)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range sortingTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			actual := MergeSort(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v to give: %v, actual: %v", tt.input, tt.expected, actual)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range sortingTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			actual := QuickSort(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v to give: %v, actual: %v", tt.input, tt.expected, actual)
			}
		})
	}
}

func sortingBenchmark() []int {
	return []int{30, 85, 45, 53, 4, 10, 44, 33, 97, 3, 34, 74, 15, 7, 40, 73, 46, 82, 32, 63, 62, 25, 50, 76, 2, 21, 5, 47, 8, 65, 52, 39, 96, 66, 89, 67, 58, 68, 14, 57, 92, 38, 17, 35, 60, 18, 48, 83, 27, 13, 72, 28, 86, 70, 20, 59, 95, 69, 9, 75, 87, 94, 81, 42, 98, 80, 26, 22, 88, 78, 90, 61, 41, 77, 16, 1, 84, 11, 51, 100, 99, 6, 79, 49, 43, 93, 64, 54, 91, 55, 29, 37, 71, 19, 56, 12, 23, 24, 36, 31}
}

// avoid optimising it away
var benchmarkValue interface{}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkValue = BubbleSort(sortingBenchmark())
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkValue = MergeSort(sortingBenchmark())
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkValue = QuickSort(sortingBenchmark())
	}
}
