package main

import (
	"sort"

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

func isComplete(g *simple.UndirectedGraph) bool {
	// show that all nodes are reachable from each other
	nodes := graph.NodesOf(g.Nodes())
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].ID() < nodes[j].ID()
	})

	for _, n := range nodes {
		// fmt.Println(n.ID(), graph.NodesOf(k4.From(n.ID())))
		// fmt.Println("append", append(graph.NodesOf(k4.From(n.ID())), n))
		reachable := append(graph.NodesOf(g.From(n.ID())), n)
		sort.Slice(reachable, func(i, j int) bool {
			return reachable[i].ID() < reachable[j].ID()
		})

		for i := range reachable {
			if reachable[i] != nodes[i] {
				return false
			}
		}
	}

	return true
}
