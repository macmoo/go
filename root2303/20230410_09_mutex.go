package main

import (
	"fmt"
	"sync"
	"time"
)

// 패키지 전역 변수 뮤텍스
var mutex sync.Mutex

type Account struct {
	Balance int
}

func main() {
	fmt.Println("---------------------------------")
	var wg sync.WaitGroup

	account := &Account{0} // 0원 잔고 통장
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	// 뮤텍스획득
	mutex.Lock() // 뮤텍스를 확보할 때까지 대기
	// 이하 로직은 뮤텍스를 확보한
	// 단 하나의 고루틴만 실행하게 됩니다.
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

// ---------------------------------
// panic: Balance should not be negative value: -1000

// goroutine 14 [running]:
// main.DepositAndWithdraw(0xc000018098?)
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:32 +0xa8
// main.main.func1()
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:22 +0x25
// created by main.main
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:20 +0x93
// panic: Balance should not be negative value: -4000

// goroutine 11 [running]:
// main.DepositAndWithdraw(0xc000018098?)
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:32 +0xa8
// main.main.func1()
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:22 +0x25
// created by main.main
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:20 +0x93
// panic: Balance should not be negative value: -2000

// goroutine 6 [running]:
// main.DepositAndWithdraw(0xc000018098?)
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:32 +0xa8
// main.main.func1()
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:22 +0x25
// created by main.main
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:20 +0x93
// panic: Balance should not be negative value: -3000

// goroutine 10 [running]:
// main.DepositAndWithdraw(0xc000018098?)
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:32 +0xa8
// main.main.func1()
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:22 +0x25
// created by main.main
//         d:/200.dev/__git_repo/go/root2303/20230410_08_concur_trouble.go:20 +0x93
// exit status 2
