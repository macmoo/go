package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for { // 무한루프
		fmt.Println("입력해라")
		var number int
		_, err := fmt.Scanln(&number) // 한줄입력
		if err != nil {
			fmt.Println("숫자 입력해라 :")
			stdin.ReadString('\n') // 키보드 버퍼 클리어
			continue
		}
		fmt.Printf("입력한 숫자는 %d이다.\n", number)
		if number%2 == 0 { // 짝수 검사
			break
		}
	}
	fmt.Println("for문 종료")
}
