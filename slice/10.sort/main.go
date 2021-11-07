package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	slice := []int{6, 3, 1, 5, 4, 2}

	fmt.Println(slice)
	sort.Ints(slice)
	fmt.Println(slice)

	students := []Student{
		{"c", 31},
		{"a", 20},
		{"b", 21},
		{"d", 19},
	}

	fmt.Println(students)
	sort.Sort(Students(students))
	fmt.Println(students)
}
