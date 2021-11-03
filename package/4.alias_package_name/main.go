package main

import (
	"fmt"
	htmlplate "html/template"
	"text/template"
)

func main() {
	fmt.Println(template.New("foo").Parse(`{{define "T"}}Hello{{end}}`))
	fmt.Println(htmlplate.New("foo").Parse(`{{define "T"}}Hello{{end}}`))
}
