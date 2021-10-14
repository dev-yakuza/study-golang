package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{10, 20, 30, 40, 50}

	fmt.Println(a)
	fmt.Println(b)

	b = a
	fmt.Println(a)
	fmt.Println(b)

	// ERROR
	// c := [3]int{1, 2, 3}
	// d := [5]int{10, 20, 30, 40, 50}

	// d = c
	// fmt.Println(c)
	// fmt.Println(d)

	// ERROR
	// c := [...]int{1, 2, 3}
	// d := [...]int{10, 20, 30, 40, 50}

	// d = c
	// fmt.Println(c)
	// fmt.Println(d)
}
