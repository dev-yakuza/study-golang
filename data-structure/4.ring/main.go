package main

import (
	"container/ring"
	"fmt"
)

func main() {
	ring := ring.New(5)

	n := ring.Len()

	for i := 0; i < n; i++ {
		ring.Value = i
		ring = ring.Next()
	}

	fmt.Println("From Front")
	for i := 0; i < n; i++ {
		fmt.Print(ring.Value, " ")
		ring = ring.Next()
	}

	fmt.Println("")
	fmt.Println("From Back")
	for i := 0; i < n; i++ {
		fmt.Print(ring.Value, " ")
		ring = ring.Prev()
	}
}
