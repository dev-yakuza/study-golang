package main

import "fmt"

// func Divide(a, b int) (int, bool) {
// 	if b == 0 {
// 		return 0, false
// 	}

// 	return a / b, true
// }

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}

	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(10, 2)
	fmt.Println(c, success)
	d, success := Divide(10, 0)

	fmt.Println(d, success)
}
