package gostructures

import (
	"fmt"
	"reflect"
	"testing"
)

func Hard04(array []uint, n uint) uint {
	sum := uint(0)

	for i := uint(1); i <= n; i++ {
		sum ^= i
	}

	b := uint(0)
	for _, i := range array {
		b ^= i
	}
	return sum ^ b
}

func TestHard04(t *testing.T) {
	var out uint

	out = Hard04([]uint{7, 4, 1, 2, 3, 6}, 7)
	if out != 5 {
		t.Errorf("should be 5, got: %d", out)
	}

	out = Hard04([]uint{1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15}, 15)
	if out != 7 {
		t.Errorf("should be 7, got: %d", out)
	}
}

func Hard07(counts map[string]uint, synonyms [][2]string) map[string]uint {
	// map canonical names to synonyms
	synonymMap := make(map[string]string)

	for _, pair := range synonyms {
		canon, synonym := pair[0], pair[1]

		if name, ok := synonymMap[canon]; ok {
			canon = name
		} else if name, ok := synonymMap[synonym]; ok {
			canon, synonym = name, canon
		} else {
			synonymMap[canon] = canon
		}
		synonymMap[synonym] = canon

		count := counts[synonym]
		delete(counts, synonym)
		counts[canon] += count
	}

	return counts
}

func TestHard07(t *testing.T) {
	names := map[string]uint{
		"John":        15,
		"Jon":         12,
		"Chris":       13,
		"Kris":        4,
		"Christopher": 19,
	}

	synonyms := [][2]string{
		{"Jon", "John"},
		{"John", "Johnny"},
		{"Chris", "Kris"},
		{"Christopher", "Chris"},
	}

	expected := map[string]uint{
		"Jon":   27,
		"Chris": 36,
	}

	actual := Hard07(names, synonyms)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("combined names are wrong\nexpected: %v\nactual:   %v", expected, actual)
	}
}

func IsHoppable(array []uint) bool {
	if len(array) == 0 {
		return true
	}
	if len(array) == 1 {
		return array[0] > 0
	}

	// go back from 1 to n steps from the end
	// if the height is >= i, then we can hop from there
	for i := 1; i <= len(array); i++ {
		candidate := len(array) - i
		if array[candidate] >= uint(i) {
			if IsHoppable(array[:candidate]) {
				//	fmt.Println("is hoppable:", array[:candidate])
				return true
			}
		}
	}
	return false
}

func TestIsHoppable(t *testing.T) {
	array := []uint{4, 2, 0, 0, 2, 0}
	if !IsHoppable(array) {
		t.Errorf("should be hoppable")
	}

	// add another 0, so its no longer hoppable
	array = append(array, 0)
	if IsHoppable(array) {
		t.Errorf("should not be hoppable")
	}

	// when it can't even start
	array = []uint{4, 3, 2, 1, 0, 5, 5, 5, 5}
	if IsHoppable(array) {
		t.Errorf("should not be hoppable")
	}

	// when it can only finish right from the start
	array = []uint{8, 0, 0, 0, 0, 0, 0, 0}
	if !IsHoppable(array) {
		t.Errorf("should be hoppable")
	}

	// but add another, and again its not hoppable
	array = append(array, 0)
	if IsHoppable(array) {
		t.Errorf("should not be hoppable")
	}
}

func LongestConsecutiveCharacter(chars string) (rune, int) {
	maxChar := '?'
	maxCount := 0
	char := '?'
	count := 0

	for _, c := range chars {
		if c == char {
			count++
		} else {
			char = c
			count = 1
		}
		if count > maxCount {
			maxChar = char
			maxCount = count
		}
	}

	return maxChar, maxCount
}

func TestLongestConsecutiveCharacter(t *testing.T) {
	var char rune
	var length int

	char, length = LongestConsecutiveCharacter("A")
	if char != 'A' || length != 1 {
		t.Errorf("expected: A, 1, actual: %v, %d", char, length)
	}

	char, length = LongestConsecutiveCharacter("AAAAA")
	if char != 'A' || length != 5 {
		t.Errorf("expected: A, 5, actual: %v, %d", char, length)
	}

	char, length = LongestConsecutiveCharacter("AABCDDBBBEA")
	if char != 'B' || length != 3 {
		t.Errorf("expected: B, 3, actual: %v, %d", char, length)
	}
}

func LongestCommonSubsequence(p, q []rune) []rune {
	cache := make(map[[2]string][]rune)

	return longestCommonSubsequence(p, q, &cache)
}

func longestCommonSubsequence(p, q []rune, cache *map[[2]string][]rune) []rune {
	cacheKey := [2]string{string(p), string(q)}
	if cached, ok := (*cache)[cacheKey]; ok {
		return cached
	}

	if len(p) == 0 || len(q) == 0 {
		return []rune{}
	} else if p[len(p)-1] == q[len(q)-1] {
		result := longestCommonSubsequence(p[:len(p)-1], q[:len(q)-1], cache)
		result = append(result, p[len(p)-1])

		(*cache)[cacheKey] = result
		return result
	} else {
		result1 := longestCommonSubsequence(p[:len(p)], q[:len(q)-1], cache)
		result2 := longestCommonSubsequence(p[:len(p)-1], q[:len(q)], cache)

		if len(result2) > len(result1) {
			result1 = result2
		}

		(*cache)[cacheKey] = result1
		return result1
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	var p, q, expected, actual []rune

	p = []rune("BATD")
	q = []rune("ABACD")
	expected = []rune("BAD")

	actual = LongestCommonSubsequence(p, q)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", string(expected), string(actual))
	}

	p = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	q = []rune("ZYXWVUTSRQPONMLKJIHGFEDCBA")
	expected = []rune("Z")

	actual = LongestCommonSubsequence(p, q)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v, actual: %v", string(expected), string(actual))
	}
}

// given a list of integers,
// we need to choose a sublist which consists of no consecutive elements
// such that the sum of elements is maximised
//
// eg.
// 	 { 30, 15, 60, 75, 45, 15, 15, 45 }
// maximised:
//   { 30,     60,     45,         45 }
//
func Hard16(array []int) int {
	cache := make(map[int]int)
	return hard16(array, 0, &cache)
}

func hard16(array []int, index int, cache *map[int]int) int {
	if index >= len(array) {
		return 0
	}

	if cached, ok := (*cache)[index]; ok {
		return cached
	}

	// including the first element or not
	with := array[index] + hard16(array, index+2, cache)
	without := hard16(array, index+1, cache)

	if with > without {
		(*cache)[index] = with
		return with
	} else {
		(*cache)[index] = without
		return without
	}
}

func Hard16Iterative(array []int) int {
	best := make([]int, len(array)+2) // buffer with some trailing 0s

	for i := len(array) - 1; i >= 0; i-- {
		// take this element and skip the next one
		with := array[i] + best[i+2]

		// don't take this element
		without := best[i+1]

		if with > without {
			best[i] = with
		} else {
			best[i] = without
		}
	}
	return best[0]
}

func TestHard16(t *testing.T) {
	testcases := []struct {
		in, out []int
		sum     int
	}{
		{
			in:  []int{30, 15, 60, 75, 45, 15, 15, 45},
			out: []int{30, 60, 45, 45},
			sum: 180,
		},

		{
			in:  []int{75, 105, 120, 75, 90, 135},
			out: []int{75, 120, 135},
			sum: 330,
		},
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			actual1 := Hard16(tt.in)
			if actual1 != tt.sum {
				t.Errorf("Hard16 returned the wrong sum")
			}

			actual2 := Hard16Iterative(tt.in)
			if actual2 != tt.sum {
				t.Errorf("Hard16Iterative returned the wrong sum")
			}
		})
	}
}
