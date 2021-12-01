package main

import (
	"fmt"
)

func firstCall() {
	fmt.Println("(2) firstCall function start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in firstCall", r)
		}
	}()

	group()
	fmt.Println("(2) firstCall function end")
}

func group() {
	fmt.Println("(3) group function start")
	fmt.Printf("4/2 = %d\n", divide(4, 2))
	fmt.Printf("4/0 = %d\n", divide(4, 0))
	fmt.Println("(3) group function end")
}

func divide(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

func main() {
	fmt.Println("(1) main function start")
	firstCall()
	fmt.Println("(1) main function end")
}
