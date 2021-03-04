package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
)

func main() {

	c4 := makeCycle(4)
	// k4 := makeComplete(4)
	// k6 := makeComplete(6)

	fmt.Println("c4", graph.EdgesOf(c4.Edges()), graph.NodesOf(c4.Nodes()))
	// fmt.Println("k4", graph.NodesOf(k4.Nodes()), graph.NodesOf(k4.Nodes()))
	// fmt.Println("k6", graph.NodesOf(k6.Nodes()))

}
