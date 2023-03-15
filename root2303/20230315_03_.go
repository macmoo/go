package main

import "fmt"

type User struct {
	name string
	age  int
}

func newUser(name string, age int) *User {
	var u = User{name, age}
	return &u // 메모리에서 사라지지 않음
}

func main() {
	userPointer := newUser("AAA", 22)
	fmt.Printf("name=%s, age=%d\n", userPointer.name, userPointer.age)
}

// name=AAA, age=22
