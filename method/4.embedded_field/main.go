package main

import "fmt"

type Member struct {
	Name string
	Age  int
}

func (p *Member) AddAge() {
	p.Age += 1
}

type Group struct {
	Member
	Grade int
}

func main() {
	g := Group{Member: Member{Name: "John", Age: 20}, Grade: 1}
	fmt.Println(g)
	g.AddAge()
	fmt.Println(g)
}
