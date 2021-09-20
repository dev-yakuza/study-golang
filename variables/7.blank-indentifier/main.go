package main

import "fmt"

func main() {
	tmp := make(map[string]string)

	tmp["name"] = "John"
	tmp["job"] = "Programmer"

	fmt.Println(tmp)

	for k, v := range tmp {
		fmt.Println(k, v)
	}

	for k := range tmp {
		fmt.Println(k)
	}

	for _, v := range tmp {
		fmt.Println(v)
	}
}
