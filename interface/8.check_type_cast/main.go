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

func main() {
	p := &Person{}
	s := Student(p)
	s.Learn()

	t, ok := s.(Teacher)
	fmt.Println(ok)
	if ok {
		t.Teach()
	}

	if t, ok := s.(Teacher); ok {
		t.Teach()
	}
}
