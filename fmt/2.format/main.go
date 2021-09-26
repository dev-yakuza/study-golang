package main

import "fmt"

func main() {
	a := 1
	b := 10
	c := 100
	d := 1000
	e := 10000

	fmt.Printf("a: %5d\n", a)
	fmt.Printf("b: %5d\n", b)
	fmt.Printf("c: %5d\n", c)
	fmt.Printf("d: %5d\n", d)
	fmt.Printf("e: %5d\n", e)

	fmt.Println()
	fmt.Printf("a: %05d\n", a)
	fmt.Printf("b: %05d\n", b)
	fmt.Printf("c: %05d\n", c)
	fmt.Printf("d: %05d\n", d)
	fmt.Printf("e: %05d\n", e)

	fmt.Println()
	fmt.Printf("a: %-5d\n", a)
	fmt.Printf("b: %-5d\n", b)
	fmt.Printf("c: %-5d\n", c)
	fmt.Printf("d: %-5d\n", d)
	fmt.Printf("e: %-5d\n", e)
}
