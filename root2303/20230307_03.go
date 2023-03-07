package main

import "fmt"

func main() {
	a := 3              // int
	var b float64 = 3.5 // float64

	// var c int = b	// error
	var c int = int(b)  // 형변환
	d := float64(a * c) // int -> float64

	var e int64 = 7
	f := int64(d) * e // float64 -> int64

	var g int = int(b * 3)
	var h int = int(b) * 3

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g, h)
}

// 3 3.5 3 9 7 63
// 10 9
