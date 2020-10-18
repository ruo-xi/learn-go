package main

import "fmt"

func main() {
	s := "宇大帅比"
	rs := []rune(s)
	fmt.Println(s, "\n", rs)
	rs[1] = '太'
	rs[3] = '了'
	fmt.Println(string(rs))

	s1 := "hello"
	fmt.Println([]byte(s1), "\n", []rune(s1))

}
