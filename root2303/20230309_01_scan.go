package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	var d int
	var e int
	var f int

	fmt.Print("------------------------------\n")
	n, err := fmt.Scan(&a, &b)
	if err != nil {
		// 에러 발생시 에러 출력
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
	fmt.Print("------------------------------\n")
	m, err2 := fmt.Scanf("%d %d\n", &c, &d)
	if err2 != nil {
		fmt.Print(m, err2)
	} else {
		fmt.Print(m, c, d)
	}
	fmt.Print("------------------------------\n")
	l, err3 := fmt.Scanln(&e, &f)
	if err3 != nil {
		fmt.Print(l, err3)
	} else {
		fmt.Print(l, e, f)
	}
	fmt.Print("------------------------------\n")

}
