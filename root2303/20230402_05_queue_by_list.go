package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue)Push(val interface{}){
	q.v.PushBack(val)
}

func (q *Queue)Pop() interface{]){
	front := q.
}

func main() {

	fmt.Println("---------------------------------")
}
