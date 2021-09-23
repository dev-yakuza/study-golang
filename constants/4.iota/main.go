package main

import "fmt"

// const (
// 	Red   int = iota
// 	Blue  int = iota
// 	Green int = iota
// )

// const (
// 	Red int = iota
// 	Blue
// 	Green
// )

const (
	Red = iota
	Blue
	Green
)

func main() {
	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)
}
