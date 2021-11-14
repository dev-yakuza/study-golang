package main

import "fmt"

type SampleInterface interface {
	SampleMethod()
}

func main() {
	var s SampleInterface
	fmt.Println(s)
	s.SampleMethod()
}
