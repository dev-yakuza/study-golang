package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiple(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return multiple
	} else {
		return nil
	}
}

func main() {
	var op func(int, int) int

	op = getOperator("+")
	result := op(10, 20)
	fmt.Println(result)

	op = getOperator("*")
	result = op(10, 20)
	fmt.Println(result)
}
