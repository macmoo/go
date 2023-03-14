package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1바이트
	B int  // 8바이트
	C int8
	D int
	E int8
}
type User2 struct {
	A int
	B int
	C int8
	D int8
	E int8
}
type User3 struct {
	C int8
	D int8
	E int8
	A int
	B int
}

func main() {
	user := User{1, 2, 3, 4, 5}
	user2 := User2{1, 2, 3, 4, 5}
	user3 := User3{1, 2, 3, 4, 5}
	fmt.Println("Size:", unsafe.Sizeof(user))
	fmt.Println("Size:", unsafe.Sizeof(user2))
	fmt.Println("Size:", unsafe.Sizeof(user3))

}
