package main

import "fmt"

const (
	C1 int = iota
	C2 int = iota
	C3 int = iota
)

const (
	D1 uint = iota + 1
	D2
	D3
)

const (
	E1 uint = 1 << iota
	E2
	E3
)

func main() {
	fmt.Println(C1) // 0
	fmt.Println(C2) // 1
	fmt.Println(C3) // 2
	fmt.Println("---------------")
	fmt.Println(D1) // 1
	fmt.Println(D2) // 2
	fmt.Println(D3) // 3
	fmt.Println("---------------")
	fmt.Println(E1) // 1
	fmt.Println(E2) // 2
	fmt.Println(E3) // 4
	fmt.Println("---------------")
}

// ---------------
// 0
// 1
// 2
// ---------------
// 1
// 2
// 3
// ---------------
// 1
// 2
// 4
// ---------------
