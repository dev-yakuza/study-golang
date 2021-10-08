package main

import "fmt"

func main() {
	switch a := 2; a {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 6, 8:
		fmt.Println("Even")
	default:
		fmt.Println("Please insert 0 < value < 10.")
	}
}
