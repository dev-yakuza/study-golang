package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("(%s) Try to eat\n", name)
		first.Lock()
		fmt.Printf("(%s) Grab %s\n", name, firstName)
		second.Lock()
		fmt.Printf("(%s) Grab %s\n", name, secondName)
		fmt.Printf("(%s) Eating\n", name)
		time.Sleep(time.Duration(rand.Int()) * time.Millisecond)
		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "Fork", "Spoon")
	go diningProblem("B", spoon, fork, "Spoon", "Fork")

	wg.Wait()
}
