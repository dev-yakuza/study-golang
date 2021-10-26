package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello, World!"
	str2 := str1

	strHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	strHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(strHeader1)
	fmt.Println(strHeader2)
}
