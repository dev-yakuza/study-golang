package main

import "fmt"

func main() {
	var p *int

	fmt.Println(p)
	if p != nil {
		fmt.Println("Assigned")
	}
}
