package server

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})
	return mux
}

func main() {
	fmt.Println("---------------------------------")
	http.ListenAndServe(":3000", MakeWebHandler())
}
