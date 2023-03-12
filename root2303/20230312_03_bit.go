package main

import "fmt"

func main() {
	var a byte = 2
	var b byte = 10
	fmt.Println("-----------------------")
	fmt.Printf("%08b\n", a)
	fmt.Println("-----------------------")
	fmt.Printf("%08b\n", ^a)
	fmt.Println("-----------------------")
	// &^ 비트클리어 연산자
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", ^a)
	fmt.Printf("%08b\n", b&^a)
	fmt.Println("-----------------------")
}
