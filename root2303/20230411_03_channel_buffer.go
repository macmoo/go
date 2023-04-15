package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("square:%d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
func main() {
	fmt.Println("---------------------------------")
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	// for range에서 데이터를 모두 처리하고 난 다음에 채널이 닫힌 상태이면
	// for문을 종료해서 프로그램이 정상 종료될 수 있습니다.
	close(ch) // 개선판) 채널닫음
	wg.Wait()
}

// ---------------------------------
// square:0
// square:4
// square:16
// square:36
// square:64
// square:100
// square:144
// square:196
// square:256
// square:324
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc000107f48?)
//         C:/Program Files/Go/src/runtime/sema.go:62 +0x27
// sync.(*WaitGroup).Wait(0x836a98?)
//         C:/Program Files/Go/src/sync/waitgroup.go:116 +0x4b
// main.main()
//         d:/200.dev/__git_repo/go/root2303/20230411_03_channel_buffer.go:27 +0x11d

// goroutine 6 [chan receive]:
// main.square(0x0?, 0x0?)
//         d:/200.dev/__git_repo/go/root2303/20230411_03_channel_buffer.go:10 +0xa5
// created by main.main
//         d:/200.dev/__git_repo/go/root2303/20230411_03_channel_buffer.go:22 +0xe5
// exit status 2

// 개선판
// ---------------------------------
// square:0
// square:4
// square:16
// square:36
// square:64
// square:100
// square:144
// square:196
// square:256
// square:324
