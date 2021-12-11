package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "v", 3)

	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("v"); v != nil {
		n := v.(int)
		fmt.Println("Square:", n*n)
	}
	wg.Done()
}
