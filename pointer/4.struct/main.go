package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func main() {
	var data Data
	var p1 *Data = &data
	var p2 *Data = &Data{}

	fmt.Println(*p1)
	fmt.Println(*p2)
}
