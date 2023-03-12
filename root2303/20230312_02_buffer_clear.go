package main

import (
	"bufio" // io
	"fmt"
	"os" // 표준입출력
)

func main() {
	// 표준입력을 읽는 객체
	stdin := bufio.NewReader(os.Stdin)
	var a int
	var b int
	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		/// 표준 입력스트림 지우기
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
