package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Printf("---------------------------------\n")
	s1 := "안녕하세요. 한글 문자열입니다."
	s2 := s1
	fmt.Printf(s1)
	fmt.Printf("\n")
	fmt.Printf(s2)
	fmt.Printf("---------------------------------\n")
	fmt.Printf("s1 addr : %p\n", &s1)
	fmt.Printf("s2 addr : %p\n", &s2)
	fmt.Printf("---------------------------------\n")
	header_s1 := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	header_s2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Println(header_s1)
	fmt.Println(header_s2)
	fmt.Printf("---------------------------------\n")

}
