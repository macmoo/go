package main

import "fmt"

func main() {

	fmt.Println("---------------------------------")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[0:3]
	fmt.Println("s1:", s1, len(s1), cap(s1))
	fmt.Println("s2:", s2, len(s2), cap(s2))
	fmt.Println("---------------------------------")
	s3 := s1[:3]
	fmt.Println("s3:", s3, len(s3), cap(s3))
	fmt.Println("---------------------------------")
	s4 := s1[2:len(s1)]
	s5 := s1[2:]
	fmt.Println("s4:", s4, len(s4), cap(s4))
	fmt.Println("s5:", s5, len(s5), cap(s5))
	fmt.Println("---------------------------------")
	s6 := s1[:]
	fmt.Println("s6:", s6, len(s6), cap(s6))
	fmt.Println("---------------------------------")
	s7 := s1[1:3:4]
	fmt.Println("s7:", s7, len(s7), cap(s7))
	// 인덱스 1부터 인덱스 2까지 집어냅니다. 그래서 [2, 3]이 됩니다.
	// cap은 최대인덱스가 4이므로 4에서 시작인덱스값인 1을 뺀 값인 3이 됩니다.
	fmt.Println("---------------------------------")
}

// ---------------------------------
// s1: [1 2 3 4 5] 5 5
// s2: [1 2 3] 3 5
// ---------------------------------
// s3: [1 2 3] 3 5
// ---------------------------------
// s4: [3 4 5] 3 3
// s5: [3 4 5] 3 3
// ---------------------------------
// s6: [1 2 3 4 5] 5 5
// ---------------------------------
// s7: [2 3] 2 3
// ---------------------------------
