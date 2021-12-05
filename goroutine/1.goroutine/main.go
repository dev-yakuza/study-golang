package main

import (
	"fmt"
	"time"
)

func PrintAlphabet() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, v := range alphabet {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	PrintAlphabet()
	fmt.Println("")

	PrintNumbers()
	fmt.Println("")

	go PrintAlphabet()
	go PrintNumbers()
	time.Sleep(3 * time.Second)
}
