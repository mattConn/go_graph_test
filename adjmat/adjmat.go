package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// matrix
	const dim = 3
	data := make([]float64, dim*dim)
	for i := range data {
		data[i] = float64(i)
	}
	m := mat.NewSymDense(dim, data)
	fmt.Println(m)

}
