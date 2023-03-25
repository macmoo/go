package main

import (
	"fedex"
	"fmt"
)

func SendBok(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	fmt.Println("---------------------------------")
	sender := &fedex.FedexSender{}
	SendBook("어린 완자", sender)
	SendBook("그리스인 조르바", sender)
}
