package main

import "fmt"

func main() {
	a := 6
	b := 5

	fmt.Printf("Result: %08b\n", a&b)
	fmt.Printf("Result: %08b\n", a|b)
	fmt.Printf("Result: %08b\n", a^b)
	fmt.Printf("Result: %08b\n", a&^b)
}
