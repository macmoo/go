package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	// account 구조체의 모든 값이 메서드 호출 시와 메서드 결괏값 반환 시 두 번 복사됩니다.
	// mainB는 메서드 호출 이후 변경된 값을 갖는 새로운 객체를 나타냅니다.
	// a3, mainA, mainB 모두 다른 메모리를 주소를 갖는 서로 다른 객체입니다.
	a3.balance -= amount
	return a3
}

func main() {
	fmt.Println("---------------------------------")
	var mainA *account = &account{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance) // 70
	fmt.Println("---------------------------------")
	mainA.withdrawValue(20)
	fmt.Println(mainA.balance) // 70
	fmt.Println("---------------------------------")
	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance) // 50
	fmt.Println("---------------------------------")
	mainB.withdrawPointer(30)  // 20
	fmt.Println(mainB.balance) // 50
	fmt.Println("---------------------------------")

	// mainA.withdrawValue(20)에서 mainA는 *account 포인터 변수이고
	// withdrawValue()는 account 값 타입을 리시버로 받는 메서드입니다. 포인터인
	// mainA로 바로 호출할 수 없고 (*mainA).withdrawValue(20)과 같이 값 타입으로 변
	// 환하여 호출하여야 합니다. 하지만 Go 언어에서는 이럴 때는 자동으로 mainA의 값으로
	// 변환하여 호출합니다.
	// 비슷하게 mainB.withdrawPointer(30)에서 mainB는 account 값 타입 변수이고
	// withdrawPointer() 메서드는 *account 포인터를 리시버로 받는 메서드입니다. 역시
	// 값 타입인 mainB로 바로 호출할 수 없고, (&mainB).withdrawPointer(30)와 같이 주
	// 소 연산자를 사용해서 포인터로 변환 후에 호출해야 합니다. 하지만 Go 언어에서는 자동
	// 으로 mainB의 메모리 주솟값으로 변환하여 호출합니다.
	// 포인터 메서드는 인스턴스 중심이고 값 타입 메서드는 값 중심이 됩니다.
}

// ---------------------------------
// 70
// ---------------------------------
// 70
// ---------------------------------
// 50
// ---------------------------------
// 20
// ---------------------------------
