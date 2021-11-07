package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	deleteIdx := 2

	fmt.Println(slice)
	for i := deleteIdx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	slice = []int{1, 2, 3, 4, 5, 6}
	deleteIdx = 2

	fmt.Println(slice)
	slice = append(slice[:deleteIdx], slice[deleteIdx+1:]...)
	fmt.Println(slice)

	slice = []int{1, 2, 3, 4, 5, 6}
	deleteIdx = 2

	fmt.Println(slice)
	copy(slice[deleteIdx:], slice[deleteIdx+1:])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
