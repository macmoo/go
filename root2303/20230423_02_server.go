package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("---------------------------------")
	// ServeMux인스턴스 생성
	mux := http.NewServeMux()
	// 인스턴스에 핸들러 등록
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})
	// mux인스턴스 사용
	http.ListenAndServe(":3000", mux)
}
