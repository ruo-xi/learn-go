package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	t2 := time.Date(2008, 7, 15, 16, 30, 28, 0, time.Local)
	fmt.Println(t2)

	s1 := t1.Format("2006年1月2日 3:04:05")
	fmt.Println(s1)

	s2 := "2009年10月10日"
	t3, err := time.Parse("2006年1月2日", s2)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(t3)

	t1.Date()
	t1.Clock()
}
