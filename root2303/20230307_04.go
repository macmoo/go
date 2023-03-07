package main

import "fmt"

var g int = 10 // 전역변수

func main() {
	var m int = 20 // 지역변수

	{
		var s int = 50 // 지역변수
		fmt.Println(m, s, g)
	}

	// m = s + 20 // error
}
