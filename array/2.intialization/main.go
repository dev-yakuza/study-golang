package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	days := [3]string{"Mon", "Tue", "Wed"}
	fmt.Println(days)

	e := [5]int{1, 2}
	fmt.Println(e)

	f := [5]int{1: 10, 3: 30}
	fmt.Println(f)

	g := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", g)
	fmt.Println(g)

	h := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", h)
	fmt.Println(h)
}
