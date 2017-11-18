package gostructures

import "testing"

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap()
	heap.Insert(10)

	if heap.Max() != 10 {
		t.Errorf("max should be 10, actual: %d", heap.Max())
	}

	heap.Insert(5)
	heap.Insert(1)
	heap.Insert(20)
	heap.Insert(17)
	heap.Insert(3)

	if heap.Max() != 20 {
		t.Errorf("max should be 20, actual: %d", heap.Max())
	}

	heap.Pop()
	if heap.Max() != 17 {
		t.Errorf("max should be 17, actual: %d", heap.Max())
	}

	heap.Pop()
	if heap.Max() != 10 {
		t.Errorf("max should be 10, actual: %d", heap.Max())
	}

	heap.Pop()
	if heap.Max() != 5 {
		t.Errorf("max should be 5, actual: %d", heap.Max())
	}

	heap.Pop()
	if heap.Max() != 3 {
		t.Errorf("max should be 3, actual: %d", heap.Max())
	}

	heap.Pop()
	if heap.Max() != 1 {
		t.Errorf("max should be 1, actual: %d", heap.Max())
	}

	if len(heap.tree) != 1 {
		t.Error("not clearing the tree correctly")
	}
}
