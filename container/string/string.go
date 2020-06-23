package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "h你好"
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	println()
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	println()

	fmt.Println("rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}
	println()

	fmt.Println(strings.Fields(" 你好 世界 "))
	fmt.Println(strings.Split("你好 世界 宇宙", " "))
	fmt.Println(strings.Join([]string{"你好", "世界"}, " "))

	println(strings.Contains(s, "h"))

}
