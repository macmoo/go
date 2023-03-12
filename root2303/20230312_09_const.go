package main

import "fmt"

func main() {
	const C int = 10
	var b int = C * 20
	// C = 20 에러
	fmt.Println(b)
	// fmt.Println(&C) // 에러 상수주소 출력 불가능
}
