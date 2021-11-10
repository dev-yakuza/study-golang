package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (r Person) greeting() string {
	return "Hello, " + r.Name
}

func main() {
	p := Person{Name: "Rob", Age: 4}
	fmt.Println(p.greeting())
}
