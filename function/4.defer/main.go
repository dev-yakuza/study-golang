// package main

// import "fmt"

// func print() {
// 	defer fmt.Println("World!")
// 	fmt.Println("Hello")
// }

// func main() {
// 	print()
// }

package main

import "os"

func readFile() {
	f, err := os.Open("file.txt")

	if err != nil {
		return
	}

	f.Close()
}

func main() {
	readFile()
}
