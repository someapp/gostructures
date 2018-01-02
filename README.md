GoStructures
============

[![Build Status](https://travis-ci.org/matthewrudy/gostructures.svg?branch=master)](https://travis-ci.org/matthewrudy/gostructures)

Building Classic Data Structures in Go

[Stack](stack.go)
-----------------

A Stack implemented as a linked list

``` go
stack := NewStack()
stack.Push(23)

value, err := stack.Pop()
assert(value == 23)

value, err = stack.Pop()
assert(err == errors.New("empty stack"))
```

[MultiStack](multi_stack.go)
----------------------------

A Stack of Stacks
that resizes

``` go
// with a maxlength of 64 for each stack
stack := NewMultiStack(64)
stack.Push(23)

value, err := stack.Pop()
assert(value == 23)
```

[Queue](queue.go)
-----------------

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

[BitVector](bit_vector.go)
--------------------------

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

[MaxHeap](max_heap.go)
----------------------

A tree structure that always keeps the max value at the top

``` go
heap := NewMaxHeap()
heap.Push(7)
heap.Push(23)

assert(heap.Max() == 23)

heap.Pop()
assert(heap.Max() == 7)
```

[Graph](graph.go)
-----------------

A directional Graph implemented with a vertex list
with adjency list for each vertex

``` go
graph := NewGraph()
graph.AddVertex("a")
graph.AddVertex("b")
graph.AddVertex("c")

graph.AddEdge("a", "b")
graph.AddEdge("a", "c")

assert(GraphConnected(&graph, "a", "b"))
assert(!GraphConnected(&graph, "b", "c"))
```

[Trie](trie.go)
---------------

A trie implemented as a tree

``` go
trie := NewTrie()
trie.AddWord("test")

assert(trie.ContainsWord("test"))
assert(!trie.ContainsWord("tes"))

assert(trie.WordsWithPrefix("tes") == []string{"test"})
```
