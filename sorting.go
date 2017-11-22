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
	if len(array) < 2 {
		return array
	}

	// pivot on the first element
	pivot := array[0]

	left := make([]int, 0, len(array)) // we'll use this as our return
	right := make([]int, 0, len(array)-1)

	for _, element := range array[1:] {
		if element < pivot {
			left = append(left, element)
		} else {
			right = append(right, element)
		}
	}

	leftSorted := QuickSort(left)
	copy(left, leftSorted)
	left = append(left, pivot)
	left = append(left, QuickSort(right)...)

	return left
}
