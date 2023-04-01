package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

// Student 타입의 String() 메서드
func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	// *Student 타입으로 타입 변환
	s := stringer.(*Student)
	fmt.Printf("Age:%d\n", s.Age)
}

func main() {
	// *Student 타입 변수 s 선언 및 초기화
	s := &Student{15}
	// 변수 s를 인터페이스 인수로 PrintAge() 함수 호출
	PrintAge(s)
	fmt.Println("---------------------------------")
	fmt.Println(s.String())
	fmt.Println("---------------------------------")
}

// ---------------------------------
// Age:15
// ---------------------------------
// Student Age:15
// ---------------------------------
