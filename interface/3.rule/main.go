package main

import "fmt"

type SampleInterface interface {
	String() string
	String(a int) string // Error: duplicated method name
	_(x int) int         // Error: no name method
}

func main() {
	var s SampleInterface
	fmt.Println(s)
}
