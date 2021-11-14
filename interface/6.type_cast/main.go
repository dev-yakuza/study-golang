package main

import "fmt"

type Human interface {
	Learn()
}

type Teacher struct {
	Name string
}

func (m *Teacher) Learn() {
	fmt.Println("Teacher can learn")
}

func (m *Teacher) Teach() {
	fmt.Println("Teacher can teach")
}

type Student struct {
	Name string
}

func (m *Student) Learn() {
	fmt.Println("Student can learn")
}

func Study(h Human) {
	if h != nil {
		h.Learn()

		var s *Student
		s = h.(*Student)
		fmt.Printf("Name: %v\n", s.Name)
		// s.Teach() // ERROR
	}
}

func main() {
	Study(&Student{Name: "John"})
}

// Example 2
/*
package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

func main() {
	var stringer Stringer
	s := stringer.(*Student)
	fmt.Println(s)
}
*/
