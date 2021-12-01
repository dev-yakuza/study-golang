package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')

	return line, nil
}

func WriteFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintf(file, data)
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "Hello, World!\n")
		if err != nil {
			fmt.Println("Failed to create file!", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("Failed to read file!", err)
			return
		}
	}
	fmt.Println("File contents: ", line)
}
