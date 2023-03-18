package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var char rune = '한'
	fmt.Printf("%T\n", char) // 1 char 타입 출력
	fmt.Println(char)        // 2 char값 출력
	fmt.Printf("%c\n", char) // 3 문자 출력
	fmt.Println("-------------------------------")
	str1 := "가나다라마"
	str2 := "abded"
	fmt.Printf("len(str1):%d\n", len(str1)) // UTF-8에서 한글 3바이트
	fmt.Printf("len(str2):%d\n", len(str2)) // UTF-8에서 영어 1바이트
	fmt.Printf("Sizeof(str1):%d\n", unsafe.Sizeof(str1))
	fmt.Printf("Sizeof(str2):%d\n", unsafe.Sizeof(str2))
	fmt.Println("-------------------------------")
	str3 := "Hello World"
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	fmt.Println(str3)
	fmt.Println(string(runes))
	fmt.Printf("len(str3):%d\n", len(str3))
	fmt.Printf("len(string(runes)):%d\n", len(string(runes)))
	fmt.Printf("Sizeof(str3):%d\n", unsafe.Sizeof(str3))
	fmt.Printf("Sizeof(runes):%d\n", unsafe.Sizeof(runes))

	fmt.Printf("type:%T\n", str3)
	fmt.Printf("type:%T\n", string(runes))
}

// int32
// 54620
// 한
// -------------------------------
// len(str1):15
// len(str2):5
// Sizeof(str1):16
// Sizeof(str2):16
// -------------------------------
// Hello World
// Hello World
// len(str3):11
// len(string(runes)):11
// Sizeof(str3):16
// Sizeof(runes):24
// type:string
// type:string
