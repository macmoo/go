package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일수 없다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	fmt.Println("---------------------------------")
	divide(9, 3)
	divide(9, 0)
}

// ---------------------------------
// 9 / 3 = 3
// panic: b는 0일수 없다.

// goroutine 1 [running]:
// main.divide(0x9?, 0x3?)
//         d:/200.dev/__git_repo/go/root2303/20230410_04_panic.go:7 +0x105
// main.main()
//         d:/200.dev/__git_repo/go/root2303/20230410_04_panic.go:15 +0x71
// exit status 2
// ---------------------------------
