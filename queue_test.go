package gostructures

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue()

	if queue.Length() != 0 {
		t.Errorf("empty queue should have length 0, actual:", queue.Length())
	}

	queue.Push(21)
	queue.Push(42)
	queue.Push(73)
	queue.Push(14)
	queue.Push(65)
	queue.Push(36)
	queue.Push(57)

	if queue.Length() != 7 {
		t.Errorf("queue should have length 7, actual:", queue.Length())
	}

	top, err := queue.Pop()
	if err != nil {
		t.Errorf("should be no error popping a queue")
	}
	if top != 21 {
		t.Errorf("popped element is wrong, should be 21, actual:", top)
	}

	if queue.Length() != 6 {
		t.Errorf("queue should have length 6, actual:", queue.Length())
	}

	queue.Pop() // 6
	queue.Pop() // 5
	queue.Pop() // 4
	queue.Pop() // 3
	queue.Pop() // 2

	top, err = queue.Pop()

	if err != nil {
		t.Errorf("should be no error popping a stack")
	}
	if top != 57 {
		t.Errorf("popped element is wrong, should be 57, actual:", top)
	}

	if queue.Length() != 0 {
		t.Errorf("empty queue should have length 0, actual:", queue.Length())
	}

	top, err = queue.Pop()
	if err == nil {
		t.Errorf("should be an error popping an empty queue")
	}
}
