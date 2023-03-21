package main

import "fmt"

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}
	slice1 := array1[1:2]
	fmt.Println("---------------------------------")
	fmt.Println("array1:", array1)
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("---------------------------------")
	array1[1] = 100
	fmt.Println("array1:", array1)
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("---------------------------------")
	slice1 = append(slice1, 500)
	fmt.Println("array1:", array1)
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("---------------------------------")
}

// ---------------------------------
// array1: [1 2 3 4 5]
// slice1: [2] 1 4
// ---------------------------------
// array1: [1 100 3 4 5]
// slice1: [100] 1 4
// ---------------------------------
// array1: [1 100 500 4 5]
// slice1: [100 500] 2 4
// ---------------------------------
