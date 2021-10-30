package main

import (
	"fmt"

	"github.com/dev-yakuza/study-golang/module/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}
