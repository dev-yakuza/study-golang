package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Name  string
	Class int
}

type Memory struct {
	A int8
	B int
	C int8
	D int
	E int8
}

type MemoryAlignment struct {
	A int8
	C int8
	E int8
	B int
	D int
}

func main() {
	s := Student{"John", 1}
	var str string = "John"
	var i int = 1

	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(str))
	fmt.Println(unsafe.Sizeof(i))

	m := Memory{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(m))

	ma := MemoryAlignment{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(ma))
}
