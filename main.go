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
	k100 := makeComplete(100)
	c100 := makeCycle(100)

	fmt.Println("k4 is complete:", isComplete(k4))
	fmt.Println("c4 is complete: ", isComplete(c4))
	fmt.Println("k100 is complete: ", isComplete(k100))
	fmt.Println("c100 is complete: ", isComplete(c100))

}
