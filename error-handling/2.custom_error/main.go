package main

import (
	"errors"
	"fmt"
)

func errorTest(i int) (int, error) {
	switch i {
	case 0:
		return i, fmt.Errorf("Error: %d", i)
	case 1:
		return i, errors.New("Error")
	default:
		return i, nil
	}
}

func main() {
	i, err := errorTest(0)
	if err != nil {
		fmt.Println(err)
	}

	i, err = errorTest(1)
	if err != nil {
		fmt.Println(err)
	}

	i, err = errorTest(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
