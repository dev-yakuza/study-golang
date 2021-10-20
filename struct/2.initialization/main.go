package main

import "fmt"

type Student struct {
	Name  string
	Class int
	No    int
}

func main() {
	var s Student
	fmt.Println(s)

	var s1 Student = Student{"Tom", 1, 2}
	var s2 Student = Student{
		"John",
		1,
		3,
	}
	fmt.Println(s1)
	fmt.Println(s2)

	var s3 Student = Student{Name: "Deku", Class: 1, No: 3}
	fmt.Println(s3)
}
