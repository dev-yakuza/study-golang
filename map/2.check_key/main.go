package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = ""
	fmt.Printf("name: %s / %T\n", m["name"], m["name"])
	fmt.Printf("country: %s / %T\n", m["country"], m["country"])

	v, ok := m["name"]
	fmt.Println(v, ok)

	v, ok = m["country"]
	fmt.Println(v, ok)
}
