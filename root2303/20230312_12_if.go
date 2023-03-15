package main

import "fmt"

func getMyAge() (int, bool) {
	return 40, true
}
func main() {

	// if 초기문; 조건문 {}
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("Youne", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("BEAUTIFUL", age)
	} else {
		fmt.Println("ERROR")
	}
	// fmt.Println("You're age is ", age) // error : age 는 소멸됨
}

// BEAUTIFUL 40
