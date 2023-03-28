package main

import "fmt"

func PrintVal(v interface{}) {
	switch t:=v.type(){
	case int:
		fmt.Printf("v is int %d\n", int(t))
	}
}

func main() {
	fmt.Println("---------------------------------")
}
