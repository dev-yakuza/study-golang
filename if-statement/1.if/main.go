package main

import "fmt"

func main() {
	v := 89

	if v > 90 {
		fmt.Println("B")
	} else if v > 80 {
		fmt.Println("B")
	} else if v > 70 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}
}
