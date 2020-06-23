package main

import "fmt"

func main() {
	var s []int //zero value for slice is nil
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
		printSlice(s)

	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	printSlice(s3)

	copy(s2, s1)
	printSlice(s2)

	//s2[:3] + s2[4:]
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	front := s2[0]
	s2 = s2[1:]
	printSlice(s2)
	println(front)

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	printSlice(s2)
	println(tail)
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}
