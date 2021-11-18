package main

import "fmt"

type PrintConsole func(string)

func PrintHello(p PrintConsole) {
	p("Hello, World!")
}

func main() {
	PrintHello(func(msg string) {
		fmt.Println(msg)
	})
}
