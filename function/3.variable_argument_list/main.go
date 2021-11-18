package main

import "fmt"

func sum(nums ...int) int {
	fmt.Printf("nums type: %T\n", nums)

	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}
