package main

import "fmt"

type Stringer interface {
	String() string
}
type Student struct{}
type Actor struct{}

func (s *Student) String() string { return "Student" }
func (a *Actor) String() string   { return "Actor" }

func ConvertType(stringer Stringer) {
	student := stringer.(*Student)
	fmt.Println(student.String())
}

func main() {
	fmt.Println("---------------------------------")
	// actor := &Actor{}
	student := &Student{}

	// 에러
	// panic: interface conversion: main.Stringer is *main.Actor, not *main.Student
	// ConvertType(actor)
	ConvertType(student)
	fmt.Println("---------------------------------")
}

// ---------------------------------
// Student
// ---------------------------------
