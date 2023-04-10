// 작업 분할 방식과 역할 분할 방식으로 뮤텍스 없이 동시성 프로그래밍을 할 수 있습니다.
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업시작\n", j.index) // 각 작업
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

// 각 고루틴은 할당된 작업만 하므로 고루틴 간 간섭이 발생하지 않습니다.
// 그래서 뮤텍스가 필요없습니다.
// 예를 들면 100개 파일을 읽어서 분석할 때는 파일별로 고루틴을 할당해서 수행하면 되겠죠.
// 중요한 점은 고루틴 간의 간섭을 없애는 것.
func main() {
	fmt.Println("---------------------------------")
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		job := jobList[i] // 각 작업을 고루틴으로 실행
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}

// ---------------------------------
// 9 작업시작
// 4 작업시작
// 0 작업시작
// 1 작업시작
// 2 작업시작
// 3 작업시작
// 6 작업시작
// 5 작업시작
// 7 작업시작
// 8 작업시작
// 8 작업 완료 - 결과: 64
// 9 작업 완료 - 결과: 81
// 4 작업 완료 - 결과: 16
// 0 작업 완료 - 결과: 0
// 3 작업 완료 - 결과: 9
// 6 작업 완료 - 결과: 36
// 1 작업 완료 - 결과: 1
// 2 작업 완료 - 결과: 4
// 5 작업 완료 - 결과: 25
// 7 작업 완료 - 결과: 49
// ---------------------------------
