package greeting

import (
	"fmt"

	"github.com/dev-yakuza/study-golang/package/name"
)

func Print() {
	fmt.Println("Hello, ", name.Name)
}
