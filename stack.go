package gostructures

import "errors"

type Stack struct {
	head   *stackNode
	length int
}

func NewStack() Stack {
	return Stack{}
}

type stackNode struct {
	next  *stackNode
	value interface{}
}

func (s *Stack) Push(value interface{}) {
	next := stackNode{
		next:  s.head,
		value: value,
	}

	s.head = &next
	s.length++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.length == 0 {
		return 0, errors.New("empty stack")
	}
	head := s.head
	s.head = head.next
	s.length--

	return head.value, nil
}

func (s *Stack) Length() int {
	return s.length
}
