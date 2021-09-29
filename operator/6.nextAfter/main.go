package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f\n", c)
	fmt.Printf("%0.18f\n", a+b)
	fmt.Printf("%0.18f == %0.18f (%v)\n", c, a+b, c == a+b)
	fmt.Printf("%0.18f == %0.18f (%v)\n", c, a+b, c == math.Nextafter(a+b, c))
}
