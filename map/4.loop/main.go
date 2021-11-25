package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[19] = Product{Name: "TV", Price: 3000}
	m[16] = Product{Name: "Phone", Price: 1000}
	m[18] = Product{Name: "PC", Price: 500}
	m[17] = Product{Name: "Tablet", Price: 2000}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
