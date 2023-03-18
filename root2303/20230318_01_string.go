package main

import "fmt"

func main() {
	fmt.Printf("---------------------------------\n")
	s1 := "Hello"
	s2 := "Hell"
	s3 := "Hello"

	fmt.Printf("%s == %s : %v\n", s1, s2, s1 == s2)
	fmt.Printf("%s != %s : %v\n", s1, s2, s1 != s2)
	fmt.Printf("%s == %s : %v\n", s1, s3, s1 == s3)
	fmt.Printf("%s != %s : %v\n", s1, s3, s1 != s3)
	fmt.Printf("---------------------------------\n")
}

// ---------------------------------
// Hello == Hell : false
// Hello != Hell : true
// Hello == Hello : true
// Hello != Hello : false
