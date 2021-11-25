package main

import "fmt"

func main() {
	m := make(map[int]int)

	m[1] = 0
	m[2] = 2
	m[3] = 3

	v, ok := m[1]
	fmt.Println(v, ok)

	delete(m, 1)

	v, ok = m[1]
	fmt.Println(v, ok)
}
