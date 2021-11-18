package main

import "fmt"

type OperateFunc func(int, int) int

func getOperator(op string) OperateFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
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
