package gostructures

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()

	if stack.Length() != 0 {
		t.Errorf("empty stack should have length 0, actual:", stack.Length())
	}

	stack.Push(21)
	stack.Push(42)
	stack.Push(73)
	stack.Push(14)
	stack.Push(65)
	stack.Push(36)
	stack.Push(57)

	if stack.Length() != 7 {
		t.Errorf("empty stack should have length 7, actual:", stack.Length())
	}

	top, err := stack.Pop()
	if err != nil {
		t.Errorf("should be no error popping a stack")
	}
	if top != 57 {
		t.Errorf("popped element is wrong, should be 57, actual:", top)
	}

	if stack.Length() != 6 {
		t.Errorf("empty stack should have length 6, actual:", stack.Length())
	}

	stack.Pop() // 6
	stack.Pop() // 5
	stack.Pop() // 4
	stack.Pop() // 3
	stack.Pop() // 2

	top, err = stack.Pop()

	if err != nil {
		t.Errorf("should be no error popping a stack")
	}
	if top != 21 {
		t.Errorf("popped element is wrong, should be 21, actual:", top)
	}

	if stack.Length() != 0 {
		t.Errorf("empty stack should have length 0, actual:", stack.Length())
	}

	top, err = stack.Pop()
	if err == nil {
		t.Errorf("should be an error popping an empty stack")
	}
}
