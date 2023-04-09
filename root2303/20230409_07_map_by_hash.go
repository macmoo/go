package main

import "fmt"

const M = 10 // 분모

// 해시함수
func hash(d int) int {
	return d % M // 나머지 연산
}

func main() {
	m := [M]int{} // 배열 생성

	m[hash(23)] = 10
	m[hash(259)] = 50
	fmt.Println("---------------------------------")
	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	fmt.Printf("%d = %d\n", 259, m[hash(259)])
	fmt.Println("---------------------------------")
	// 해시충돌
	m[hash(33)] = 99
	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	fmt.Printf("%d = %d\n", 33, m[hash(33)])
	fmt.Println("---------------------------------")
	// 해시충돌의 해결방법은 인덱스 위치마다 값이 아니라 리스트를 저장하는것.
	fmt.Println("---------------------------------")
}
