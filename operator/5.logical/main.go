package main

import "fmt"

func main() {
	a := true
	b := false

	fmt.Println("Result:", a && b)
	fmt.Println("Result:", a || b)
	fmt.Println("Result:", !a)
}
