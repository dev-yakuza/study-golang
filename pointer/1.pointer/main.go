package main

import "fmt"

func main() {
	var a int = 10
	var p *int

	fmt.Println(a)

	p = &a
	fmt.Printf("%v\n", &a)
	fmt.Printf("%v\n", p)

	*p = 20
	fmt.Println(a)
	fmt.Println(*p)
}
