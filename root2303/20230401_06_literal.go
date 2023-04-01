package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		// 함수 리터럴을 사용해서 더하기 함수를 정의하고 반환
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		// 함수 리터럴을 사용해서 곱하기 함수를 정의하고 반환
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}
func main() {
	fn := getOperator("*")
	// 함수 타입 변수를 사용해서 함수 호출
	result := fn(3, 4)
	fmt.Println("---------------------------------")
	fmt.Println(result)
	fmt.Println("---------------------------------")
	fmt.Printf("Type:%T\n", fn)
	fmt.Println("---------------------------------")
	result = func(a, b int) int {
		return a * b
	}(10, 5)
	fmt.Println(result)
	fmt.Println("---------------------------------")
}

// ---------------------------------
// 12
// ---------------------------------
// Type:main.opFunc
// ---------------------------------
// 50
// ---------------------------------
