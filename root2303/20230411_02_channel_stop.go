package main

import "fmt"

func main() {
	fmt.Println("---------------------------------")
	ch := make(chan int) // 크기가 0인 채널 생성
	// 다른 데서 채널의 데이터를 빼지 않기 때문에
	// main()가 여기서 멈춤
	ch <- 9
	fmt.Println("Never Print") // 실행되지 않음
}

// ---------------------------------
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
//         d:/200.dev/__git_repo/go/root2303/20230411_02_channel_stop.go:8 +0x71
// exit status 2
