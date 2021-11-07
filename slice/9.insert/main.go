package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	insertIdx := 2

	fmt.Println(slice)
	slice = append(slice, 0)
	for i := len(slice) - 2; i >= insertIdx; i-- {
		slice[i+1] = slice[i]
	}
	slice[insertIdx] = 100
	fmt.Println(slice)

	slice = []int{1, 2, 3, 4, 5, 6}
	insertIdx = 2

	fmt.Println(slice)
	slice = append(slice[:insertIdx], append([]int{100}, slice[insertIdx:]...)...)
	fmt.Println(slice)

	slice = []int{1, 2, 3, 4, 5, 6}
	insertIdx = 2

	fmt.Println(slice)
	slice = append(slice, 0)
	copy(slice[insertIdx+1:], slice[insertIdx:])
	slice[insertIdx] = 100
	fmt.Println(slice)
}
