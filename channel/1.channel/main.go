package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	time.Sleep(time.Second)

	ch <- 3
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	fmt.Println("Square start")
	n := <-ch
	fmt.Println("Square:", n*n)
	wg.Done()
	fmt.Println("Square end")
}
