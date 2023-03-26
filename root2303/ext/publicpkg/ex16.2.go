package main

import (
	"fmt"

	"ch16/ex16.2/publicpkg"
)

func main() {
	fmt.Printf("---------------------------------\n")
	fmt.Println("PI:", publicpkg.PI)
	publicpkg.PublicFunc()
	fmt.Printf("---------------------------------\n")
	var myint publicpkg.MyInt = 10
	fmt.Println("MyInt:", myint)
	fmt.Printf("---------------------------------\n")
	var mystruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("mystruct:", mystruct)
	fmt.Printf("---------------------------------\n")
	publicpkg.privateFunc()
	fmt.Printf("---------------------------------\n")
}

// ---------------------------------
// PI: 3.1415
// This is a public function 100
// ---------------------------------
// MyInt: 10
// ---------------------------------
// mystruct: {18 }
// ---------------------------------
