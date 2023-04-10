package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("---------------------------------")
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저") // A는 포크 먼저
	go diningProblem("B", spoon, fork, "수저", "포크") // B는 수저 먼저

	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock() // 첫 번째 뮤텍스를 획득 시도

		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock() // 두 번째 뮤텍스를 획득 시도

		fmt.Printf("%s %s 획득\n", name, secondName)
		fmt.Printf("%s 밥을 먹습니다.\n", name)

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock() // 뮤텍스 반납
		first.Unlock()
	}
	wg.Done()
}

// ---------------------------------
// B 밥을 먹으려 합니다.
// B 수저 획득
// B 포크 획득
// B 밥을 먹습니다.
// A 밥을 먹으려 합니다.
// A 포크 획득
// A 수저 획득
// A 밥을 먹습니다.
// B 밥을 먹으려 합니다.
// A 밥을 먹으려 합니다.
// A 포크 획득
// B 수저 획득
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc000008090?)
//         C:/Program Files/Go/src/runtime/sema.go:62 +0x27
// sync.(*WaitGroup).Wait(0xe399e8?)
//         C:/Program Files/Go/src/sync/waitgroup.go:116 +0x4b
// main.main()
//         d:/200.dev/__git_repo/go/root2303/20230410_10_deadlock.go:23 +0x1b1

// goroutine 6 [sync.Mutex.Lock]:
// sync.runtime_SemacquireMutex(0x10?, 0x60?, 0x2?)
//         C:/Program Files/Go/src/runtime/sema.go:77 +0x26
// sync.(*Mutex).lockSlow(0xc0000180c0)
//         C:/Program Files/Go/src/sync/mutex.go:171 +0x165
// sync.(*Mutex).Lock(...)
//         C:/Program Files/Go/src/sync/mutex.go:90
// main.diningProblem({0xe18aa0, 0x1}, 0xc000018098, 0xc0000180c0, {0xe19087, 0x6}, {0xe19081, 0x6})
//         d:/200.dev/__git_repo/go/root2303/20230410_10_deadlock.go:32 +0x1f1
// created by main.main
//         d:/200.dev/__git_repo/go/root2303/20230410_10_deadlock.go:20 +0x14f

// goroutine 7 [sync.Mutex.Lock]:
// sync.runtime_SemacquireMutex(0x10?, 0x60?, 0x2?)
//         C:/Program Files/Go/src/runtime/sema.go:77 +0x26
// sync.(*Mutex).lockSlow(0xc000018098)
//         C:/Program Files/Go/src/sync/mutex.go:171 +0x165
// sync.(*Mutex).Lock(...)
//         C:/Program Files/Go/src/sync/mutex.go:90
// main.diningProblem({0xe386b0, 0x1}, 0xc0000180c0, 0xc000018098, {0xe19081, 0x6}, {0xe19087, 0x6})
//         d:/200.dev/__git_repo/go/root2303/20230410_10_deadlock.go:32 +0x1f1
// created by main.main
//         d:/200.dev/__git_repo/go/root2303/20230410_10_deadlock.go:21 +0x1a5
// exit status 2

// D:\200.dev\__git_repo\go>
