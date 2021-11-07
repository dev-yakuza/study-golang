package main

import "fmt"

func main() {
	slice1 := make([]int, 10)
	fmt.Println(slice1)

	slice2 := make([]int, 2, 10)
	fmt.Println(slice2)
}
