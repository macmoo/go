package main

import (
	"container/ring"
	"fmt"
)

func main() {
	fmt.Println("---------------------------------")
	r := ring.New(5)
	n := r.Len()
	fmt.Printf("ring size:%d\n", n)
	fmt.Println("---------------------------------")

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}
	fmt.Println("\n---------------------------------")

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
	fmt.Println("\n---------------------------------")
}

// ---------------------------------
// ring size:5
// ---------------------------------
// A B C D E
// ---------------------------------
// A E D C B
// ---------------------------------
