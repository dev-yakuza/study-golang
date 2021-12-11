package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	go PrintTick(ctx)

	time.Sleep(time.Second * 5)
	cancel()

	wg.Wait()
}

func PrintTick(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done:", ctx.Err())
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
