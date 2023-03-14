package main

import "fmt"

type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이:%d 번호:%d 점수:%.2f\n", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{15, 23, 88.2}

	student2 := student

	PrintStudent(student2)
	student.No = 99
	PrintStudent(student2)
	PrintStudent(student)
}

// 나이:15 번호:23 점수:88.20
// 나이:15 번호:23 점수:88.20
// 나이:15 번호:99 점수:88.20
