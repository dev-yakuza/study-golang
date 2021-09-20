package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int32 = 10

	// fmt.Println(num1 + num2)
	fmt.Println(num1 + int(num2))

	var num3 float32 = 10.32
	fmt.Println(num1 + int(num3))
}
