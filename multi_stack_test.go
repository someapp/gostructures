package gostructures

import (
	"testing"
)

func TestMultStack(t *testing.T) {
	stack := NewMultiStack(2)

	if stack.Length() != 0 {
		t.Errorf("empty stack should have length 0, actual: %d", stack.Length())
	}

	stack.Push(21)
	stack.Push(42)
	stack.Push(73)
	stack.Push(14)
	stack.Push(65)
	stack.Push(36)
	stack.Push(57)

	if stack.Length() != 7 {
		t.Errorf("empty stack should have length 7, actual: %d", stack.Length())
	}

	// implementation test
	if len(stack.stacks) != 4 {
		t.Errorf("should use 4 individual stacks, actual: %d", len(stack.stacks))
	}

	top, err := stack.Pop()
	if err != nil {
		t.Errorf("should be no error popping a stack")
	}
	if top != 57 {
		t.Errorf("popped element is wrong, should be 57, actual: %d", top)
	}

	if stack.Length() != 6 {
		t.Errorf("empty stack should have length 6, actual: %d", stack.Length())
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
		t.Errorf("popped element is wrong, should be 21, actual: %d", top)
	}

	if stack.Length() != 0 {
		t.Errorf("empty stack should have length 0, actual: %d", stack.Length())
	}

	// implementation test
	if len(stack.stacks) != 0 {
		t.Errorf("should use 1 individual stack when empty, actual: %d", len(stack.stacks))
	}

	top, err = stack.Pop()
	if err == nil {
		t.Errorf("should be an error popping an empty stack")
	}
}
