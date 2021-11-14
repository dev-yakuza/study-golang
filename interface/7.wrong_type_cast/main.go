package main

import "fmt"

type Teacher interface {
	Teach()
}

type Student interface {
	Learn()
}

type Person struct {
}

func (f *Person) Learn() {
	fmt.Println("Person can learn")
}

func Teach(s Student) {
	t := s.(Teacher)
	t.Teach()
}

func main() {
	p := &Person{}
	Teach(p)
}
