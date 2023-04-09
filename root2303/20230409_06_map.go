package main

import "fmt"

func PrintMap(m map[int]int) {
	for k, v := range m {
		fmt.Printf("key:%d, value:%d\n", k, v)
	}
}

func main() {
	fmt.Println("---------------------------------")
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	PrintMap(m)
	fmt.Println("---------------------------------")

	delete(m, 3) // map 삭제
	delete(m, 4)
	fmt.Println(m[3])
	fmt.Println(m[1])
	fmt.Println("---------------------------------")
	PrintMap(m)
	fmt.Println("---------------------------------")
}
