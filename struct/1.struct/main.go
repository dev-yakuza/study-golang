package main

import "fmt"

type Student struct {
	Name  string
	Class int
	No    int
}

func main() {
	var s Student
	s.Name = "Tom"
	s.Class = 1
	s.No = 1

	fmt.Println(s)
	fmt.Printf("%v\n", s)
	fmt.Printf("Name: %s, Class: %d, No: %d\n", s.Name, s.Class, s.No)
}
