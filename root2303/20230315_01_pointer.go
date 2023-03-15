package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값:%p\n", p)
	fmt.Printf("p가 가리키는 메모리 값:%d\n", *p)

	*p = 100
	fmt.Printf("p의 값:%p\n", p)
	fmt.Printf("p가 가리키는 메모리 값:%d\n", *p)
}

// p의 값:0xc00009e058
// p가 가리키는 메모리 값:500
// p의 값:0xc00009e058
// p가 가리키는 메모리 값:100
