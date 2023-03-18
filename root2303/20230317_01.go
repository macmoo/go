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
}
