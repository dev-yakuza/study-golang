package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 2, 10)

	fmt.Printf("slice1(%p): len=%d cap=%d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2(%p): len=%d cap=%d %v\n", slice2, len(slice2), cap(slice2), slice2)
}
