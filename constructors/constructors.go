package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

func main() {
	g1 := simple.NewDirectedGraph()
	g2 := simple.NewUndirectedGraph()
	MakeCompleteGraph(g1)
	MakeCompleteGraph(g2)
	fmt.Println(graph.NodesOf(g1.Nodes()))
	fmt.Println(graph.NodesOf(g2.Nodes()))

}

// MakeCompleteGraph does what it says on the tin.
func MakeCompleteGraph(g graph.Builder) {
	g.SetEdge(g.NewEdge(simple.Node(0), simple.Node(1)))
}
