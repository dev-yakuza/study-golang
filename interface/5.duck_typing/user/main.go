package user

import "fmt"

type User struct {
	Name string
}

func (u User) Walk() {
	fmt.Println("Walking")
}
func (u User) Talk() {
	fmt.Println("Talking")
}
