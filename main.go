package main

import "fmt"

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a =%v\n", a)
	s1 := a[:]
	fmt.Printf("s1=%v\n", s1)
	s2 := a[1:4]
	fmt.Printf("s2=  %v\n", s2)
	fmt.Printf("len(s2)=%d, cap(s2)=%d\n", len(s2), cap(s2))
	s3 := s2[1:5]
	fmt.Printf("s3=    %v\n", s3)
	fmt.Printf("len(s3)=%d, cap(s3)=%d\n", len(s3), cap(s3))
}
