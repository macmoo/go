package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업")
	case "wednesday", "thursday":
		fmt.Println("수, 목요일은 쉬은날")
	}
}

// 수, 목요일은 쉬은날
