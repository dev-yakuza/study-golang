package main

import "fmt"

type myInt int

func main() {
	var myNum myInt = 10
	var num int = 10

	fmt.Println(myNum, num)
	fmt.Printf("%T\n", myNum)
	fmt.Printf("%T\n", num)
}
