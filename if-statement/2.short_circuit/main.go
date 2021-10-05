package main

import "fmt"

func first(result bool) bool {
	fmt.Println("first condition is called!")
	return result
}

func second() bool {
	fmt.Println("Second condition is called!")
	return true
}

func main() {
	fmt.Println("ex 1")
	if first(false) && second() {
	}

	fmt.Println("ex 2")
	if first(true) && second() {
	}

	fmt.Println("ex 3")
	if first(true) || second() {
	}

	fmt.Println("ex 4")
	if first(false) || second() {
	}
}
