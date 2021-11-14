package main

import "fmt"

type Human interface {
	Walk() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) Walk() string {
	return fmt.Sprintf("%s can walk", s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	s := Student{Name: "John", Age: 20}
	var h Human

	h = s
	fmt.Println(h.Walk())
	// fmt.Println(h.GetAge()) // ERROR
}
