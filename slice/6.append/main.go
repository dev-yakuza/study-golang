package main

import "fmt"

func main() {
	slice1 := make([]int, 3)
	slice2 := append(slice1, 4)

	fmt.Println("[New splice]")
	fmt.Printf("slice(%p): len=%d cap=%d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2(%p): len=%d cap=%d %v\n", slice2, len(slice2), cap(slice2), slice2)

	fmt.Println("slice1 changed ========================================================")
	slice1[0] = 100
	fmt.Printf("slice(%p): len=%d cap=%d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2(%p): len=%d cap=%d %v\n", slice2, len(slice2), cap(slice2), slice2)

	fmt.Println("slice2 changed ========================================================")
	slice2[0] = 200
	fmt.Printf("slice(%p): len=%d cap=%d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2(%p): len=%d cap=%d %v\n", slice2, len(slice2), cap(slice2), slice2)

	slice1 = make([]int, 1, 3)
	slice2 = append(slice1, 4)

	fmt.Println("[Same slice]")
	fmt.Printf("slice(%p): len=%d cap=%d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2(%p): len=%d cap=%d %v\n", slice2, len(slice2), cap(slice2), slice2)

	fmt.Println("slice1 changed ========================================================")
	slice1[0] = 100
	fmt.Printf("slice(%p): len=%d cap=%d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2(%p): len=%d cap=%d %v\n", slice2, len(slice2), cap(slice2), slice2)

	fmt.Println("slice2 changed ========================================================")
	slice2[0] = 200
	fmt.Printf("slice(%p): len=%d cap=%d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("slice2(%p): len=%d cap=%d %v\n", slice2, len(slice2), cap(slice2), slice2)
}
