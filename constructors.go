package main

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

func makeCycle(n int) *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph()

	for i := 0; i < n-1; i++ {
		g.SetEdge(g.NewEdge(simple.Node(i), simple.Node(i+1)))
	}
	g.SetEdge(g.NewEdge(simple.Node(n-1), simple.Node(0)))

	return g
}

func makeComplete(n int) *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph()

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			g.SetEdge(g.NewEdge(simple.Node(i), simple.Node(j)))
		}
	}

	return g
}

// show that all nodes are reachable from each other
// number of edges in complete graph of n nodes ==
// n*(n-1)/2
func isComplete(g *simple.UndirectedGraph) bool {
	n := len(graph.NodesOf(g.Nodes()))

	return len(graph.EdgesOf(g.Edges())) == n*(n-1)/2
}
