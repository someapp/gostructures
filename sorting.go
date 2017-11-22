package gostructures

func BubbleSort(array []int) []int {
	sorted := make([]int, len(array))
	copy(sorted, array)

	for {
		bubbled := false
		for i := 0; i < len(sorted)-1; i++ {
			if sorted[i+1] < sorted[i] {
				sorted[i+1], sorted[i] = sorted[i], sorted[i+1]
				bubbled = true
			}
		}
		if !bubbled {
			break
		}
	}
	return sorted
}

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	split := len(array) / 2
	left := MergeSort(array[:split])
	right := MergeSort(array[split:])

	sorted := make([]int, 0, len(array))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			sorted = append(sorted, right...)
			return sorted
		}
		if len(right) == 0 {
			sorted = append(sorted, left...)
			return sorted
		}
		if right[0] < left[0] {
			sorted = append(sorted, right[0])
			right = right[1:]
		} else {
			sorted = append(sorted, left[0])
			left = left[1:]
		}
	}
	return sorted
}

func QuickSort(array []int) []int {
	sorted := make([]int, len(array))
	copy(sorted, array)

	quickSort(sorted)
	return sorted
}

func quickSort(array []int) {
	if len(array) < 2 {
		return
	}

	// pivot on the last element
	// just so we don't have to move it
	pivotIndex := len(array) - 1
	pivot := array[len(array)-1]

	leftIndex := 0

	for i, element := range array[:pivotIndex] {
		if element < pivot {
			if i != leftIndex {
				array[i], array[leftIndex] = array[leftIndex], array[i]
			}
			leftIndex++
		}
	}

	// swap the pivot to the middle
	array[leftIndex], array[pivotIndex] = array[pivotIndex], array[leftIndex]

	quickSort(array[:leftIndex])
	quickSort(array[leftIndex+1:])
}
