package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	fmt.Println("---------------------------------")
	a := &account{100}
	fmt.Printf("%d\n", a.balance)
	fmt.Println("---------------------------------")
	withdrawFunc(a, 30)
	fmt.Printf("%d\n", a.balance)
	fmt.Println("---------------------------------")
	a.withdrawMethod(30)
	fmt.Printf("%d\n", a.balance)
	fmt.Println("---------------------------------")
}

// ---------------------------------
// 100
// ---------------------------------
// 70
// ---------------------------------
// 40
// ---------------------------------
