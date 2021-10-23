package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func main() {
	p1 := &Data{}
	var p2 = new(Data)

	fmt.Println(*p1)
	fmt.Println(*p2)
}
