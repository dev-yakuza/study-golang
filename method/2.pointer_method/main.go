package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) AddAge() {
	p.Age += 1
}

type Temperature float64

func (t Temperature) Up(temp float64) Temperature {
	return t + Temperature(temp)
}

func main() {
	p1 := Person{Name: "Rob", Age: 4}

	fmt.Println("P1: ", p1)
	p1.AddAge()
	fmt.Println("P1: ", p1)

	t := Temperature(30.0)

	fmt.Println("T: ", t)
	t = t.Up(4.0)
	fmt.Println("T: ", t)
}
