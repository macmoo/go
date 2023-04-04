package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func (q *Queue) Size() int {
	return q.v.Len()
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	q := NewQueue()
	for i := 0; i < 5; i++ {
		q.Push(i)
	}

	fmt.Println("BEFORE::Queue Size:", q.Size())

	v := q.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = q.Pop()
	}
	fmt.Println("\nAFTER::Queue Size:", q.Size())
}

// fmt.Println("---------------------------------")
