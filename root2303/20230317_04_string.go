package main

import "fmt"

func main() {
	fmt.Println("---------------------------------")
	s1 := "Hello"
	s2 := "World"

	s3 := s1 + " " + s2
	fmt.Println(s3)

	s1 += " " + s2
	fmt.Println(s1)
}

// ---------------------------------
// Hello World
// Hello World
