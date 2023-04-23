package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	// 쿼리 인수가져오기
	values := r.URL.Query()
	// 특정키값 존재확인
	name := values.Get("name")
	if name == "" {
		name = "WORLD"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}
func main() {
	fmt.Println("---------------------------------")
	http.HandleFunc("/bar", barHandler) // "/bar"핸들러 등록
	http.ListenAndServe(":3000", nil)
}
