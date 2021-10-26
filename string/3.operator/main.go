package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "world!"
	fmt.Println(str1 + " " + str2)

	str := "Hello"
	str += " " + "world!"
	fmt.Println(str)

	str3 := "a"
	str4 := "b"
	fmt.Println(str3 == str4)
	fmt.Println(str3 != str4)
	fmt.Println(str3 < str4)
	fmt.Println(str3 > str4)
}
