package main

import "fmt"

func main() {
	str1 := "Hello, world!"

	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i])
	}

	fmt.Println()

	str2 := "Hello, 월드!"

	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}

	fmt.Println()

	str3 := "Hello, 월드!"
	arr := []rune(str3)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%c", arr[i])
	}
}
