package main

import "fmt"

func main() {
	a := 2

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
		fallthrough
	case 3:
		fmt.Println("a == 3")
	default:
		fmt.Println("a is not 1, 2 or 3")
	}
}
