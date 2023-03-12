package main

import "fmt"

// 멀티반환함수
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	} else {
		return a / b, true
	}
}

// 변수명 지정
func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	} else {
		result = a / b
		success = true
		return
	}
}
func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)

	d, success := Divide(9, 0)
	fmt.Println(d, success)

	e, success := Divide(9, 0)
	fmt.Println(e, success)
}
