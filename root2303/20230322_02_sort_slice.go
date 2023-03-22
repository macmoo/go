package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// []Student의 별칭 타입 Students
type Students []Student

// []Student의 별칭 타입인 Students를 만들고,
// Len(), Less(), Swap() 메서드를 만들어 sort.Interface를 사용할 수 있게 해줌.
func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	fmt.Println("---------------------------------")
	s1 := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("s1:", s1, len(s1), cap(s1))
	sort.Ints(s1)
	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("---------------------------------")
	// 구조체 정렬
	// s2 := Students{{"화랑", 31}, // SAME
	s2 := []Student{{"화랑", 31},
		{"백두산", 52},
		{"류", 42},
		{"켄", 38},
		{"송하나", 18}}
	fmt.Println("s2:", s2, len(s2), cap(s2))
	sort.Sort(Students(s2)) // []Student를 Students 타입으로 타입 변환
	fmt.Println("s2:", s2, len(s2), cap(s2))
	fmt.Println("---------------------------------")
}
