package main

import "fmt"

type Student struct {
	Age int
}

func Print(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Println("v is int", v)
	case float64:
		fmt.Println("v is float64", v)
	case string:
		fmt.Println("v is string", v)
	default:
		fmt.Printf("Not supported %T:%v", v, v)
	}
}

func main() {
	Print(10)
	Print(3.14)
	Print("Hello word")
	Print(Student{Age: 10})
}
