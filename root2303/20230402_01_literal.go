package main

import "fmt"

func main() {
	i := 0
	f := func() {
		i += 10
	}
	fmt.Println("---------------------------------")
	fmt.Println(i)
	fmt.Println("---------------------------------")
	i++
	fmt.Println(i)
	fmt.Println("---------------------------------")
	f()
	fmt.Println(i)
	fmt.Println("---------------------------------")
}
