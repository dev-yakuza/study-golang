package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintAlphabet() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, v := range alphabet {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c", v)
	}

	wg.Done()
}

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d", i)
	}

	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go PrintAlphabet()
	go PrintNumbers()

	wg.Wait()
}
