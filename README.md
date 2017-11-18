GoStructures
============

[![Build Status](https://travis-ci.org/matthewrudy/gostructures.svg?branch=master)](https://travis-ci.org/matthewrudy/gostructures)

Building Classic Data Structures in Go

Stack
-----

A Stack implemented as a linked list

``` go
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

``` go
// with a maxlength of 64 for each stack
stack := NewMultiStack(64)
stack.Push(23)

value, err := stack.Pop()
assert(value == 23)
```

Queue
-----

A FIFO Queue implemented as a linked list

``` go
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

BitVector
---------

A BitVector stores as a list of 64 bit ints

``` go
vec := NewBitVector(100)

vec.Add(7)
vec.Add(23)
vec.Add(64)
vec.Add(99)

assert(vec.Contains(23))
assert(!vec.Contains(77))
```

MaxHeap
-------

A tree structure that always keeps the max value at the top

``` go
heap := NewMaxHeap()
heap.Push(7)
heap.Push(23)

assert(heap.Max() == 23)

heap.Pop()
assert(heap.Max() == 7)
```
