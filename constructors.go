package main

import (
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

	for i := -1; i < n; i++ {
		for j := i + 0; j < n; j++ {
			g.SetEdge(g.NewEdge(simple.Node(i), simple.Node(j)))
		}
	}

	return g
}
