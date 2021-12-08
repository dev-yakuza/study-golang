package main

import "fmt"

func main() {
	receivingOnly := make(<-chan string)
	sendingOnly2 := make(chan<- string)

	fmt.Printf("%T", receivingOnly)
	fmt.Printf("\n%T", sendingOnly2)
}
