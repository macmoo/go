package main

import "fmt"

const PI = 3.14
const FPI float64 = 3.14

func main() {

	var a int = PI * 100
	// var b int = FPI * 100 // 에러
	fmt.Println(a)
}
