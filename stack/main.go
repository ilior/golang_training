package main

import (
	"fmt"
)

type Stack struct {
	h *stackItem
}

type stackItem struct {
	n *stackItem
	v interface{}
}

func (s *Stack) Push(i interface{}) {
	if s.h == nil {
		s.h = &stackItem{nil, i}
		return
	}

	s.h = &stackItem{s.h, i}

	return
}

func (s *Stack) Pop() (v interface{}) {
	if s.h == nil {
		return nil
	}

	v = s.h.v
	s.h, s.h.n = s.h.n, nil

	return
}

func (s *Stack) Print() {
	fmt.Println(s.h.v)
}

func main() {
	var stack Stack

	x := 2
	stack.Push(x)
	stack.Push(3)
	stack.Push("line")
	stack.Print()
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
