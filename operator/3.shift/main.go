package main

import "fmt"

func main() {
	a := 6

	fmt.Printf("Result: %08b\n", a)
	fmt.Printf("Result: %08b\n", a<<2)
	fmt.Printf("Result: %08b\n", a>>1)
}
