package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("---------------------------------")
	// 취소 가능한 컨텍스트 생성
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel() // 취소
	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

// ---------------------------------
// Tick
// Tick
// Tick
// Tick
// Tick
// ---------------------------------
