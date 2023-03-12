package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Second)
		fmt.Print(i, " ")
		i++
		if i == 3 {
			fmt.Print("\n")
			break
		}
	}
	fmt.Println("-------------------------------")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("-------------------------------")
}
