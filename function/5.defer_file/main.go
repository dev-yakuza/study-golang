package main

import "os"

func readFile() {
	f, err := os.Open("file.txt")

	if err != nil {
		return
	}

	f.Close()
}

func readFileDefer() {
	f, err := os.Open("file.txt")
	defer f.Close()

	if err != nil {
		return
	}
}

func main() {
	readFile()
	readFileDefer()
}
