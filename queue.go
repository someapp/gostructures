package gostructures

import "errors"

type Queue struct {
	head   *queueNode
	tail   *queueNode
	length int
}

func NewQueue() Queue {
	return Queue{}
}

type queueNode struct {
	next  *queueNode
	value int
}

func (q *Queue) Push(value int) {
	next := queueNode{
		value: value,
	}

	if q.tail != nil {
		q.tail.next = &next
		q.tail = &next
	} else {
		q.head = &next
		q.tail = &next
	}
	q.length++
}

func (q *Queue) Pop() (int, error) {
	if q.length == 0 {
		return 0, errors.New("empty queue")
	}
	head := q.head
	q.head = head.next
	q.length--

	return head.value, nil
}

func (q *Queue) Length() int {
	return q.length
}
