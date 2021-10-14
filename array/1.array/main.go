package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	const LEN = 5
	var b [LEN]int
	fmt.Println(b)

	// ERROR
	// var len = 5
	// var c [len]int
	// fmt.Println(c)
}
