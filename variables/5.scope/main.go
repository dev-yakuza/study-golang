package main

import "fmt"

func main() {
	num1 := 10

	{
		num2 := 20
		fmt.Println(num2)
	}

	fmt.Println(num1)
	// fmt.Println(num2)
}
