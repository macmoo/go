package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4)

	fmt.Println(slice)
	fmt.Println(slice2)
	slice2 = append(slice2, 5, 6, 7)
	fmt.Println(slice2)
	slice3 := make([]int, 3, 5)
	fmt.Println(slice3)
}

// [1 2 3]
// [1 2 3 4]
// [1 2 3 4 5 6 7]
// [0 0 0]
