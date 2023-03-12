package main

import "fmt"

// 별칭 ColorType를 선언하고 const 열거값 정의
type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)
	fmt.Println(Yellow)
	fmt.Println("--------------")
	switch color {
	case Red:
		return "RED"
	case Blue:
		return "BLUE"
	case Green:
		return "GREEN"
	case Yellow:
		return "YELLOW"
	default:
		return "UNDEFINED"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My Favorite color is ", colorToString(getMyFavoriteColor()))
}

// 0
// 1
// 2
// 3
