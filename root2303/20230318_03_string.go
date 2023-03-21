package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Printf("---------------------------------\n")
	var s1 = "Hello World"
	var slice []byte = []byte(s1) // "[]byte슬라이스"

	for i := 0; i < len(slice); i++ {
		fmt.Printf("[%d : %c]", slice[i], slice[i])
	}
	fmt.Printf("\n---------------------------------\n")
	fmt.Println(s1)
	fmt.Printf("%s\n", slice)
	fmt.Printf("---------------------------------\n")
	slice[2] = 'Q'
	fmt.Println(s1)
	fmt.Printf("%s\n", slice)

	s1_h := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	slice_h := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("s1\t:%x\n", s1_h)
	fmt.Printf("slice\t:%x\n", slice_h)
	fmt.Printf("---------------------------------\n")
}

// ---------------------------------
// [72 : H][101 : e][108 : l][108 : l][111 : o][32 :  ][87 : W][111 : o][114 : r][108 : l][100 : d]
// ---------------------------------
// Hello World
// Hello World
// ---------------------------------
// Hello World
// HeQlo World
// s1      :&{fd6eff b}
// slice   :&{c0000180c0 b}
// ---------------------------------
