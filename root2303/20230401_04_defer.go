package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("1 반드시 호출")
	defer f.Close()
	defer fmt.Println("2 파일 닫음")

	fmt.Println("3 파일에 HELLO WORLD를 쓴다.")
	fmt.Fprintln(f, "HELLO WORLD")
	// fmt.Fprint() 함수는 첫 번째 인수인 파일 핸들에 텍스트를 쓰는 함수.
	// 정확히는 첫 번째 인수는 파일 핸들이 아니라 io.Writer 인터페이스를 받기 때문에
	// io.Writer 인터페이스 메서드를 포함한 모든 객체는 Fprint() 함수를 사용할 수 있다.
}
