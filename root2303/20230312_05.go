package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	// 값 바꿈
	// 첫우변값을 처번째 좌변주소에
	// 두번째우변값을 두번째 좌변 주소에
	fmt.Println(a, b)
	// 복수대입연산자
	// 좌,우 개수를 맞춰야함
	a, b = b, a
	fmt.Println(a, b)

}
