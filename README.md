GoStructures
============

[![Build Status](https://travis-ci.org/matthewrudy/gostructures.svg?branch=master)](https://travis-ci.org/matthewrudy/gostructures)

Building Classic Data Structures in Go

Stack
-----

A Stack implemented as a linked list

```
stack := NewStack()
stack.Push(23)

value, err := stack.Pop()
assert(value == 23)

value, err = stack.Pop()
assert(err == errors.New("empty stack"))
```

MultiStack
----------

A Stack of Stacks
that resizes

```
// with a maxlength of 64 for each stack
stack := NewMultiStack(64)
stack.Push(23)

value, err := stack.Pop()
assert(value == 23)
```

Queue
-----

A FIFO Queue implemented as a linked list

```
queue := NewQueue()
queue.Push(21)
queue.Push(42)

value, err := queue.Pop()
assert(value == 21)

value, err = queue.Pop()
assert(value == 42)

value, err = queue.Pop()
assert(err == errors.New("empty queue"))
```