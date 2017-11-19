package gostructures

import (
	"errors"
)

type Graph struct {
	nodes  []graphNode
	labels map[string]*graphNode
}

type graphNode struct {
	label      string
	neighbours []*graphNode
}

func NewGraph() Graph {
	return Graph{
		nodes:  make([]graphNode, 0),
		labels: make(map[string]*graphNode),
	}
}

func (g *Graph) AddVertex(label string) {
	node := graphNode{
		label:      label,
		neighbours: make([]*graphNode, 0),
	}

	g.nodes = append(g.nodes, node)
	g.labels[label] = &node
}

func (g *Graph) AddEdge(from, to string) {
	fromNode, _ := g.getNode(from)
	toNode, _ := g.getNode(to)

	fromNode.neighbours = append(fromNode.neighbours, toNode)
}

func (g *Graph) getNode(label string) (*graphNode, error) {
	node, ok := g.labels[label]

	if !ok {
		return &graphNode{}, errors.New("no node found")
	}

	return node, nil
}

// Breadth-first search
func GraphConnected(g *Graph, from, to string) bool {
	if len(g.nodes) == 0 {
		return false
	}

	if from == to {
		return true
	}

	startNode, err := g.getNode(from)
	if err != nil {
		return false
	}

	endNode, err := g.getNode(to)
	if err != nil {
		return false
	}

	seen := make(map[*graphNode]bool)
	queue := []*graphNode{startNode}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if seen[node] {
			continue
		}
		seen[node] = true

		if node.label == endNode.label {
			return true
		}

		for _, neighbour := range node.neighbours {
			queue = append(queue, neighbour)
		}
	}

	return false
}
