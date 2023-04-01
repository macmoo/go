package main

import "fmt"

type Reader interface{ Read() }
type Closer interface{ Close() }
type File struct{}

func (f *File) Read() {}

// func ReadFile(reader Reader) {
// 	c := reader.(Closer) // 런타임에러
// 	c.Close()
// }

// 반환 가능 체크 버전
func ReadFile(reader Reader) {
	// 타입 변환 가능 여부를 체크해서 Close() 메서드를 호출합니다.
	c, ok := reader.(Closer) // 정상처리
	if ok {
		c.Close()
	}

	// 한 줄로 표현할 수 있습니다.
	// if c, ok := reader.(Closer); ok {
	// 	// ...
	// }
}

func main() {
	fmt.Println("---------------------------------")
	file := &File{}
	ReadFile(file)

}
