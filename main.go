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
	b := new([10]int)
	for i := 0; i < len(b); i++ {
		b[i] = i
	}
	fmt.Printf("b=%v\n", b)
	d := [6]int{6, 5, 4, 3, 2, 1}
	fmt.Printf("d= %v\n", d)
	d = a
	fmt.Printf("d= %v\n", d)
	d[0] = 7
	fmt.Printf("d= %v\n", d)
	fmt.Printf("a= %v\n", a)
}
