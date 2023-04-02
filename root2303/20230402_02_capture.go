package main

import "fmt"

func CaptureLoop1() {
	// 함수 리터럴 3개를 가진 슬라이스
	f := make([]func(), 3)
	fmt.Println("Value Loop1")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}
func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("Value Loop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	fmt.Println("---------------------------------")
	CaptureLoop1()
	fmt.Println("---------------------------------")
	CaptureLoop2()
	fmt.Println("---------------------------------")
}

// ---------------------------------
// Value Loop1
// 3
// 3
// 3
// ---------------------------------
// Value Loop2
// 0
// 1
// 2
// ---------------------------------
