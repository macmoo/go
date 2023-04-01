ring() 메서드
func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	// *Student 타입으로 타입 변환
	s := stringer.(*Student)
	fmt.Printf("Age:%d\n", s.Age)
}