package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func changeData(arg Data) {
	arg.value = 901
	arg.data[100] = 902
}
func changeDataPointer(arg *Data) {
	arg.value = 801
	arg.data[100] = 802
}
func main() {
	var data Data
	data.value = 101
	data.data[100] = 102

	fmt.Printf("----------------------------------------\n")
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
	fmt.Printf("----------------------------------------\n")
	changeData(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
	fmt.Printf("----------------------------------------\n")
	changeDataPointer(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
	fmt.Printf("----------------------------------------\n")
}

// ----------------------------------------
// value = 101
// data[100] = 102
// ----------------------------------------
// value = 101
// data[100] = 102
// ----------------------------------------
// value = 801
// data[100] = 802
// ----------------------------------------
