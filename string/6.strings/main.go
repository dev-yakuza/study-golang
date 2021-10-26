package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello, world!"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Contains(str, "hello"))
	fmt.Println(strings.Index(str, "w"))
	fmt.Println(strings.Compare("a", "b"))
}
