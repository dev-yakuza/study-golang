package main

import (
	"fmt"

	"github.com/dev-yakuza/study-golang/method/group"
)

func main() {
	member1 := group.Member{Name: "John", Age: 20}
	member2 := group.NewMember("Paul", 30)

	fmt.Println(member1)
	fmt.Println(member2)
}
