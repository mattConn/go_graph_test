package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph"

	"gonum.org/v1/gonum/graph/simple"
)

type myInt int

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

func main() {

	k4 := makeComplete(4)
	k6 := makeComplete(6)

	fmt.Println("k4", graph.NodesOf(k4.Nodes()), graph.NodesOf(k4.Nodes()))
	fmt.Println("k6", graph.NodesOf(k6.Nodes()))
	// fmt.Println("Union", set.UnionOfNodes(k4.Nodes(), k6.Nodes()))

}
