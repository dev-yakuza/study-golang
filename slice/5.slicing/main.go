package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Printf("array: len=%d %v\n", len(array), array)
	fmt.Printf("slice: len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	slice = []int{1, 2, 3, 4, 5}
	slice1 := slice[0:3]
	slice2 := slice[:3]

	fmt.Printf("slice1: len=%d cap=%d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2: len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)

	slice = []int{1, 2, 3, 4, 5}
	slice1 = slice[2:len(slice)]
	slice2 = slice[2:]

	fmt.Printf("slice1: len=%d cap=%d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2: len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)

	array = [5]int{1, 2, 3, 4, 5}
	slice = array[:]

	fmt.Printf("slice: len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	array1 := [100]int{1: 1, 2: 2, 99: 100}
	slice1 = array1[1:10]
	slice2 = slice1[2:99]

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = slice1[1:3:4]

	fmt.Printf("slice1: len=%d cap=%d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2: len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)
}
