package main

import "fmt"

func main() {
	str := "Hello, World!"
	str = "Hello"
	fmt.Println(str)

	// str[0] = "a"

	var str1 string = "Hello, World!"
	var slice []byte = []byte(str1)

	slice[0] = "h"[0]
	fmt.Println(str1)
	fmt.Printf("%s\n", slice)
}
