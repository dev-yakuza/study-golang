package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (q *Stack) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Stack) Pop() interface{} {
	back := q.v.Back()
	if back == nil {
		return nil
	}

	return q.v.Remove(back)
}

func main() {
	stack := NewStack()

	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	v := stack.Pop()
	for v != nil {
		fmt.Println(v)
		v = stack.Pop()
	}
}
