package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	fmt.Println("---------------------------------")
	var a myInt = 10
	fmt.Println(a.add(30))
	fmt.Println("---------------------------------")
	var b int = 20
	fmt.Println(myInt(b).add(50)) // 타입 변환:myInt 타입으로 변환 후 add() 메서드를 사용
	fmt.Println("---------------------------------")
}

// ---------------------------------
// 40
// ---------------------------------
// 70
// ---------------------------------
