package main

import "fmt"

type ClassInfo struct {
	Class int
	No    int
}

type Student struct {
	Class ClassInfo
	Name  string
}

func main() {
	var s Student = Student{
		Class: ClassInfo{Class: 1, No: 1},
		Name:  "John",
	}

	fmt.Println(s.Class.Class)
	fmt.Println(s.Class.No)
	fmt.Println(s.Name)
}
