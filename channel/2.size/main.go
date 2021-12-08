package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	ch := make(chan int, 1)

	go square()
	ch <- 9
	fmt.Println("Never Print")
}

func square() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}
