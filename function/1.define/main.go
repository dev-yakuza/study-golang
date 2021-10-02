package main

import "fmt"

// func Add(a int, b int) int {
// 	return a + b
// }

func Add(a, b int) int {
	return a + b
}

func main() {
	c := Add(1, 2)
	fmt.Println(c)
}
