package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["name"] = "John"
	m["country"] = "Korea"
	fmt.Println(m)

	m["city"] = "Seoul"
	fmt.Println(m)

	fmt.Printf("No key: %s / %T\n", m["language"], m["language"])
}
