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

func (s *Stack) Push(v interface{}) {
	s.v.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func (s *Stack) Size() int {
	return s.v.Len()
}

func main() {
	fmt.Println("---------------------------------")
	stack := NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
	}
	fmt.Printf("전 사이즈:%d\n", stack.Size())

	v := stack.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = stack.Pop()
	}
	fmt.Printf("\n후 사이즈:%d\n", stack.Size())
}

// ---------------------------------
// 전 사이즈:4
// 4 -> 3 -> 2 -> 1 ->
// 후 사이즈:0
// ---------------------------------
