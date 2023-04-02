package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()
	fmt.Println("\n---------------------------------")
	e4 := v.PushBack(4) // *
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Printf("\nType:%T\n", e4)
	fmt.Println("---------------------------------")
	e1 := v.PushBack(1) // *
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("\n---------------------------------")
	v.InsertBefore(3, e4) // *
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("\n---------------------------------")
	v.InsertAfter(2, e1) // *
	// Front() 메서드는 가장 첫 번째 요소를 반환
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("\n---------------------------------")
	// Back() 메서드는 가장 마지막 요소를 반환
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("\n---------------------------------")
}

// ---------------------------------
// 4
// Type:*list.Element
// ---------------------------------
// 4 1
// ---------------------------------
// 3 4 1
// ---------------------------------
// 3 4 1 2
// ---------------------------------
// 2 1 4 3
// ---------------------------------
