package main

import (
	"errors"
	"fmt"
)

type StrError struct {
	Msg string
}

func (e *StrError) Error() string {
	return fmt.Sprintf("Message: %s", e.Msg)
}

func msgError(msg string) error {
	return &StrError{Msg: msg}
}

func WrapError(msg string) error {
	err := msgError(msg)
	return fmt.Errorf("(Wrapping) %w", err)
}

func main() {
	err1 := msgError("Error")
	fmt.Println("[Normal Error]", err1)

	err2 := WrapError("Test Error Message")
	fmt.Println("[Wrapping Error]", err2)

	var strErr *StrError
	if errors.As(err2, &strErr) {
		fmt.Printf("[Failed] %s\n", strErr)
	}
}
