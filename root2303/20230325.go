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
	return fmt.Printf("HELLO, %d살, %s라고 한다.", s.Age, s.Name)
}

func main() {
	fmt.Println("---------------------------------")
	student := Student{"철수", 12}
	var stringer Stringer
	stringer = student
	fmt.Printf("%s\n", stringer.String())
}
