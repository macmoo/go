package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for { // 무한루프
		fmt.Println("입력해라:")
		var number int
		_, err := fmt.Scanln(&number) // 한줄입력

	}
}
