package main

import (
	"bufio"
	"fmt"
	"os"
)

const filename string = "_forTest/data.txt"

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) // 파일 열기
	if err != nil {
		return "", err
	}
	defer file.Close()
	// bufio.Reader 객체생성
	rd := bufio.NewReader(file) // 파일 내용 읽기
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

func main() {
	// 파일 읽기 시도
	line, err := ReadFile(filename)
	if err != nil {
		// 파일 생성
		err = WriteFile(filename, "This is WriteFile")
		// 에러 처리
		if err != nil {
			fmt.Println("파일생성 실패", err)
			return
		}
		// 다시 읽기 시도
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일읽기 실패", err)
			return
		}
	}
	fmt.Println("파일내용:", line)
}
