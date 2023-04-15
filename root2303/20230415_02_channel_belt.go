package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	fmt.Println("---------------------------------")
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)
	fmt.Println("Start Factory")

	wg.Add(3)
	// 고루틴 생성
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close Factory")
}

// 차체생산
func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

// 바퀴 설치
func InstallTire(tireCh chan *Car, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter Tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

// 도색
func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car:%s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}

// ---------------------------------
// Start Factory
// 3.03 Complete Car:Sports car Winter Tire Red
// 4.03 Complete Car:Sports car Winter Tire Red
// 5.04 Complete Car:Sports car Winter Tire Red
// 6.05 Complete Car:Sports car Winter Tire Red
// 7.06 Complete Car:Sports car Winter Tire Red
// 8.07 Complete Car:Sports car Winter Tire Red
// 9.08 Complete Car:Sports car Winter Tire Red
// 10.09 Complete Car:Sports car Winter Tire Red
// 11.09 Complete Car:Sports car Winter Tire Red
// 12.10 Complete Car:Sports car Winter Tire Red
// Close Factory
// ---------------------------------
