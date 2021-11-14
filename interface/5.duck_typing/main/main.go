package main

import (
	"fmt"

	"github.com/dev-yakuza/study-golang/interface/user"
)

type Human interface {
	Walk()
	Talk()
}

func main() {
	u := user.User{Name: "John"}
	fmt.Println(u)

	h := Human(u)
	fmt.Println(h)
}
