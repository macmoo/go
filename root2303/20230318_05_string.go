package main

import (
	"fmt"
	"strings"
)

func toUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

func toUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	fmt.Printf("---------------------------------\n")
	var str string = "Hello Wolrd"
	fmt.Println(toUpper1(str))
	fmt.Println(toUpper2(str))

}
