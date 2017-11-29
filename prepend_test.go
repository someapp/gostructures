package gostructures

import (
	"testing"
)

func prependBenchmarkSlice() []int {
	return []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	}
}

func BenchmarkPrepend_OneLine(b *testing.B) {
	toPrepend := prependBenchmarkSlice()
	for n := 0; n < b.N; n++ {
		benchmarkValue = append([]int{0}, toPrepend...)
	}
}

func BenchmarkPrepend_WithCapacitySplat(b *testing.B) {
	toPrepend := prependBenchmarkSlice()
	for n := 0; n < b.N; n++ {
		slice := make([]int, 0, len(toPrepend)+1)
		slice = append(slice, 0)
		benchmarkValue = append(slice, toPrepend...)
	}
}

func BenchmarkPrepend_WithCapacityLoop(b *testing.B) {
	toPrepend := prependBenchmarkSlice()
	for n := 0; n < b.N; n++ {
		slice := make([]int, 0, len(toPrepend)+1)
		slice = append(slice, 0)
		for _, element := range toPrepend {
			slice = append(slice, element)
		}
		benchmarkValue = slice
	}
}

// just for comparison
func BenchmarkPrepend_Append(b *testing.B) {
	toPrepend := prependBenchmarkSlice()
	for n := 0; n < b.N; n++ {
		benchmarkValue = append(toPrepend, 0)
	}
}
