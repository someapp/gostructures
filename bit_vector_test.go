package gostructures

import (
	"testing"
)

func TestBitVector(t *testing.T) {
	vec := NewBitVector(100)

	if vec.Count() != 0 {
		t.Errorf("vec should have count 0, actual: %d", vec.Count())
	}

	vec.Add(7)
	vec.Add(23)
	vec.Add(64)
	vec.Add(99)

	if vec.Count() != 4 {
		t.Errorf("vec should have count 4, actual: %d", vec.Count())
	}

	// add 23 again
	vec.Add(23)
	if vec.Count() != 4 {
		t.Errorf("vec should have count 4, actual: %d", vec.Count())
	}

	if vec.Contains(0) {
		t.Error("vec should not contain 0")
	}

	if !vec.Contains(7) {
		t.Error("vec should contain 7")
	}

	if !vec.Contains(23) {
		t.Error("vec should contain 7")
	}

	if vec.Contains(98) {
		t.Error("vec should not contain 98")
	}

	if !vec.Contains(99) {
		t.Error("vec should contain 99")
	}

	vec.Add(0)

	if !vec.Contains(0) {
		t.Error("vec should contain 0")
	}

	if vec.Count() != 5 {
		t.Errorf("vec should have count 5, actual: %d", vec.Count())
	}
}

func TestBitVector_lengths(t *testing.T) {
	var vec BitVector

	vec = NewBitVector(1)
	if len(vec.ints) != 1 {
		t.Error("vec should have 1 int")
	}

	vec = NewBitVector(64)
	if len(vec.ints) != 1 {
		t.Error("vec should have 1 int")
	}

	vec = NewBitVector(65)
	if len(vec.ints) != 2 {
		t.Error("vec should have 2 ints")
	}
}
