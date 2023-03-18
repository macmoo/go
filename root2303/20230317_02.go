package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := "A"
	b := 'a'

	fmt.Printf("%d\n", unsafe.Sizeof(a)) // 16
	fmt.Printf("%d\n", unsafe.Sizeof(b)) // 4

	str := "Hello 월드"
	rune := []rune(str)
	fmt.Printf("len(str) = %d\n", len(str))
	fmt.Printf("len(rune) = %d\n", len(rune))

}
