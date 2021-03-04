package main

import (
	"fmt"
)

type node int64

func (n node) ID() int64 {
	return int64(n)
}

func main() {

	k4 := makeComplete(4)
	c4 := makeCycle(4)

	fmt.Println("k4 is complete:", isComplete(k4))
	fmt.Println("c4 is complete: ", isComplete(c4))

}
