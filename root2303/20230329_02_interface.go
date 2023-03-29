package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker // 1 기본값은 nil입니다.
	att.Attack()     // 2 att가 nil이기 때문에 런 타임 에러가 발생합니다.
}

// D:\200.dev\__git_repo\go>go run "d:\200.dev\__git_repo\go\root2303\20230329_02_interface.go"
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x1bc856]

// goroutine 1 [running]:
// main.main()
//         d:/200.dev/__git_repo/go/root2303/20230329_02_interface.go:9 +0x16
// exit status 2
