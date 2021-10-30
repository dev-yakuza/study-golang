package greeting

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %s", name)
	return message
}
