package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	// 	Printf() 함수가 서식에 따라 문자열을 터미널에 출력하
	// 는 함수라면 Sprintf()는 화면에 출력하는 것이 아닌 string 타입으로 반환한다는 점이 다릅니다.
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	fmt.Println("---------------------------------")
	student := Student{"철수", 12}
	var stringer Stringer
	stringer = student
	fmt.Printf("%s\n", stringer.String())
}

// ---------------------------------
// 안녕! 나는 12살 철수라고 해
