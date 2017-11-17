package gostructures

import (
	"errors"
)

type MultiStack struct {
	stacks     []*Stack
	length     int
	stackLimit int
}

func NewMultiStack(stackLimit int) MultiStack {
	stack := NewStack()
	return MultiStack{
		stacks:     []*Stack{&stack},
		stackLimit: stackLimit,
	}
}

func (s *MultiStack) topStack() *Stack {
	return s.stacks[len(s.stacks)-1]
}

func (s *MultiStack) Push(value interface{}) {
	stack := s.topStack()

	if stack.Length() >= s.stackLimit {
		new := NewStack()
		stack = &new
		s.stacks = append(s.stacks, stack)
	}

	stack.Push(value)
	s.length++
}

func (s *MultiStack) Pop() (interface{}, error) {
	if s.length == 0 {
		return 0, errors.New("empty stack")
	}

	stack := s.topStack()
	value, _ := stack.Pop()

	if stack.Length() == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	s.length--

	return value, nil
}

func (s *MultiStack) Length() int {
	return s.length
}
