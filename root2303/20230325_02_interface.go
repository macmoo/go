package main

import (
	"fmt"

	"./ext/send/fedex"
	"./ext/send/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

// Fedex only
// func SendBook(name string, sender *fedex.FedexSender) {
// 	sender.Send(name)
// }

// 인터페이스 사용
func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	fmt.Println("---------------------------------")
	sender1 := &fedex.FedexSender{}
	SendBook("어린 완자", sender1)
	SendBook("그리스인 조르바", sender1)
	// ---------------------------------
	// Fedex sends 어린 완자 parcel
	// Fedex sends 그리스인 조르바 parcel
	fmt.Println("---------------------------------")
	sender2 := &koreaPost.PostSender{}
	SendBook("어린 왕자", sender2)
	SendBook("그리스인 조르바", sender2)
	fmt.Println("---------------------------------")
	fmt.Println("---------------------------------")
	fmt.Println("---------------------------------")

}
