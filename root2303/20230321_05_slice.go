package main

import "fmt"

func changeArr(arr [5]int) {
	arr[2] = 500
}

func changeSlice(slice []int) {
	slice[2] = 500
}
func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	sli1 := []int{1, 2, 3, 4, 5}
	fmt.Println("---------------------------------")
	fmt.Println("arr1:", arr1)
	fmt.Println("sli1:", sli1)
	changeArr(arr1)
	changeSlice(sli1)
	fmt.Println("---------------------------------")
	fmt.Println("arr1:", arr1)
	fmt.Println("sli1:", sli1)
}
