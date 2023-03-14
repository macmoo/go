// type 타입명 struct {
// 	필드명 타입
// 	...
// 	필드명 타입
// }

package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	// 초기값을 생략하면 모든 필드가 기본값으로 초기화됨
	var house House
	house.Address = "서울시 강동구"
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기:%d평\n", house.Size)
	fmt.Printf("가격:%.2f억원\n", house.Price)
	fmt.Println("타입:", house.Type)
	fmt.Println("------------------------------------")
	var h2 House = House{"부산시", 30, 2.5, "단독"}
	fmt.Println("주소:", h2.Address)
	fmt.Printf("크기:%d평\n", h2.Size)
	fmt.Printf("가격:%.2f억원\n", h2.Price)
	fmt.Println("타입:", h2.Type)
	fmt.Println("------------------------------------")
	var h3 House = House{Size: 33, Type: "아파트"} // 일부만 초기화
	fmt.Println("주소:", h3.Address)
	fmt.Printf("크기:%d평\n", h3.Size)
	fmt.Printf("가격:%.2f억원\n", h3.Price)
	fmt.Println("타입:", h3.Type)
}

// 주소: 서울시 강동구
// 크기:28평
// 가격:9.80억원
// 타입: 아파트
