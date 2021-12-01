package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("divide by zero")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(4, 2)
	divide(4, 0)
	divide(4, 3)
}
