package gostructures

import (
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
