package main

import (
	"bufio"
	"fmt"
	"os"
)

type LineInfo struct {
	lineNo int
	line   string
}
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	fmt.Println("---------------------------------")
	PrintFile("../_txt/hamlet.txt")
}

func PrintFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을수 없다.", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
