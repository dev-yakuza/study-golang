package main

import "fmt"

type OperateFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func multiple(a, b int) int {
	return a * b
}

func getOperator(op string) OperateFunc {
	if op == "+" {
		return add
	} else if op == "*" {
		return multiple
	} else {
		return nil
	}
}

func main() {
	var op OperateFunc

	op = getOperator("+")
	result := op(10, 20)
	fmt.Println(result)

	op = getOperator("*")
	result = op(10, 20)
	fmt.Println(result)
}
