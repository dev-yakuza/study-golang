package greeting

import "fmt"

const PI = 3.14

var globalVal = 0

func init() {
	fmt.Println("init()", globalVal)
	globalVal++
}

func PrintGreeting() {
	print()
}

func print() {
	fmt.Println("Hello, World!", globalVal)
}
