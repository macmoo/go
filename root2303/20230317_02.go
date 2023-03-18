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
	fmt.Printf("len(str) = %d\n", len(str)) // 바이트 길이
	fmt.Printf("len(rune) = %d\n", len(rune))
}

// len(str) = 12 바이트
// len(rune) = 8
