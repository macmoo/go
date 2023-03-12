package main

import "fmt"

func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After : ", n)
}
func main() {
	printNo(3)
}

// 3
// 2
// 1
// After :  1
// After :  2
// After :  3
