package main

import (
	"fmt"
)

type PasswordError struct {
	RequiredLen int
	Len         int
}

func (e PasswordError) Error() string {
	return "Password is too short."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{RequiredLen: 8, Len: len(password)}
	}
	// register account logic
	// ...
	return nil
}

func main() {
	err := RegisterAccount("john", "12345")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v (Required: %d Len: %d)\n", errInfo, errInfo.RequiredLen, errInfo.Len)
		}
	}
}
