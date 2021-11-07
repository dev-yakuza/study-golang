package main

import "fmt"

func changeArr(arr2 [5]int) {
	arr2[0] = 100
}
func changeSlice(slice2 []int) {
	slice2[0] = 100
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArr(arr)
	changeSlice(slice)

	fmt.Println(arr)
	fmt.Println(slice)
}
