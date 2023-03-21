package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Printf("---------------------------------\n")

	var s1 string = "Hello"
	head := (*reflect.StringHeader)(unsafe.Pointer(&s1))

	addr1 := head.Data

	s1 += " World"
	addr2 := head.Data

	s1 += " Welcome!"
	addr3 := head.Data

	fmt.Println(s1)
	fmt.Printf("addr1\t:%x\n", addr1)
	fmt.Printf("addr2\t:%x\n", addr2)
	fmt.Printf("addr3\t:%x\n", addr3)
}
