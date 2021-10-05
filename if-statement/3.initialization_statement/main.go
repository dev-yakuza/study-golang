package main

import "fmt"

func testFunc() (int, bool) {
	return 1, true
}

func main() {
	if v, success := testFunc(); success {
		fmt.Println(v)
	}

	// fmt.Println(v) // ERROR
}
