// var wg sync.WaitGroup
// wg.Add(3) // 작업 개수 설정
// wg.Done() // 작업이 완료될 때마다 호출
// wg.Wait() // 모든 작업이 완료될 때까지 대기
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // WaitGroup객체
func SumAtoB(a, b int) {
	sum := 0
	for i := 0; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지의 합계는 %d이다.\n", a, b, sum)
	wg.Done() // 작업이 완료됨을 표시
	// 각 루틴에서 SumAtoB() 함수를 완료할 때 wg.Done()을 호출되어
	// wg의 남은 작업 개수를 1씩 감소시킵니다.
}

func main() {
	fmt.Println("---------------------------------")
	wg.Add(10) // 총 작업 개수 설정(고루틴을 10개 만듬)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}
	wg.Wait() // 모든 작업이 완료되길 기다림
	// wg.Wait()를 하면 모든 작업이 완료될 때까지 종료하지 않고 대기.
	// 남은 작업 개수가 0이 되는 순간 Wait() 메서드가 끝나고 다음 줄로 넘어가게 됩니다
	fmt.Println("모든 계산이 완료됐습니다.")
}
