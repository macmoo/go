package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("---------------------------------")
	// 1. 웹 핸들러 등록
	// “/” 경로에 해당하는 HTTP 요청을 수신할때 IndexPathHandler() 함수를 호출
	// 반드시 첫 번째 인수로 http.ResponseWriter를 받고 두 번째 인수로 *http.Request 타입을 받아야 한다.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
	// 2. 웹 서버 시작
	// 첫 번째 인수로 요청을 기다릴 주소를 적어줌. 여기서는 3000번 포트에서 요청을 대기합니다.
	// 두 번째 인수로 핸들러 인스턴스를 넣어주게 되는데, nil을 넣어주면 DefaultServeMux를 사용합니다.
	// DefaultServeMux는 http.HandleFunc() 함수를 호출해 등록된 핸들러들을 사용합니다.
	http.ListenAndServe(":3000", nil)
}
