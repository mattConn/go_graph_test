package main

import (
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/graph/layout"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

func main() {

	k := makeComplete(6)

	eades := layout.EadesR2{Repulsion: 1, Rate: 0.1, Updates: 1000, Theta: 0.1, Src: rand.NewSource(1)}

	o := layout.NewOptimizerR2(k, eades.Update)

	for o.Update() {
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Add(render{o})
	p.HideAxes()

	if err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "graph.png"); err != nil {
		panic(err)
	}

}
