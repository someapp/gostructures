package gostructures

import "testing"

func TestGraph(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex("a")
	graph.AddVertex("b")
	graph.AddVertex("c")
	graph.AddVertex("d")
	graph.AddVertex("e")
	graph.AddVertex("f") // isolated

	graph.AddEdge("a", "b")
	graph.AddEdge("a", "c")
	graph.AddEdge("b", "d")
	graph.AddEdge("d", "e")

	if !GraphConnected(&graph, "a", "a") {
		t.Error("a should be connected to itself")
	}

	if !GraphConnected(&graph, "a", "b") {
		t.Error("a should be connected to b")
	}

	if !GraphConnected(&graph, "a", "e") {
		t.Error("a should be connected to e")
	}

	if GraphConnected(&graph, "a", "f") {
		t.Error("a should not be connected to f")
	}

	if GraphConnected(&graph, "b", "c") {
		t.Error("b should not be connected to c")
	}
}
