package main

import "fmt"

// PrintVal() 함수 인수로 빈 인터페이스인 interface{ }를 인수로 받습니다.
// 빈 인터페이스이기 때문에 모든 타입을 인수로 쓸 수 있습니다.
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("HELLO")
	PrintVal(Student{15})
}

// v is int 10
// v is float64 3.140000
// v is string HELLO
// Not supported type main.Student:{15}
