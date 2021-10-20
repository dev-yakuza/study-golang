package main

import "fmt"

type ClassInfo struct {
	Class int
	No    int
}

type Student struct {
	ClassInfo
	Name string
}

type DupStudent struct {
	ClassInfo
	Name string
	No   int
}

func main() {
	var s Student = Student{
		ClassInfo: ClassInfo{Class: 1, No: 1},
		Name:      "John",
	}

	fmt.Println(s.Class)
	fmt.Println(s.No)
	fmt.Println(s.Name)

	var s1 DupStudent = DupStudent{
		ClassInfo: ClassInfo{Class: 1, No: 1},
		Name:      "John",
		No:        10,
	}

	fmt.Println(s1.Class)
	fmt.Println(s1.No)
	fmt.Println(s1.Name)
	fmt.Println(s1.ClassInfo.No)
}
