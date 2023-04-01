package main

import "fmt"

type opFunc func(int, int) int

func add(a, b int) int {
	return a + b
}
func mul(a, b int) int {
	return a * b
}

// op에 따른 함수 타입 반환
func getOperataor(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	// int 타입 인수 2개를 받아서 int 타입을 반환하는 함수 타입 변수
	var operator func(int, int) int
	operator = getOperataor("*")

	// 함수 타입 변수를 사용해서 함수 호출
	var result = operator(3, 4)
	fmt.Println("---------------------------------")
	fmt.Println(result)
	fmt.Println("---------------------------------")
	var op2 opFunc
	op2 = getOperataor("+")
	result2 := op2(10, 5)
	fmt.Println(result2)
	fmt.Println("---------------------------------")
}

// ---------------------------------
// 12
// ---------------------------------
// 15
// ---------------------------------
