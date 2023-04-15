package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		// ch와 quit 양쪽을 모두 기다림
		select {
		case n := <-ch:
			fmt.Printf("square:%d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}

	}
}
func main() {
	fmt.Println("---------------------------------")
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	quit <- true
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
