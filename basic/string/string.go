package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "1\\.1"
	s2 := `1\\.1`
	fmt.Println(s1)
	fmt.Println(s2)
	//拼接
	s3 := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(s1 + s2)
	fmt.Println(s3)
	//分割
	fmt.Println(strings.Split(s2, "."))
	//包含
	fmt.Println(strings.Contains(s1, "1"))
	//前缀后缀判断
	fmt.Println(strings.HasPrefix(s1, "1"))
	fmt.Println(strings.HasSuffix(s1, "1"))
	//字串的位置
	fmt.Println(strings.LastIndex(s3, `\\`))
	//JOIN
	fmt.Println(strings.Join([]string{"1", "b"}, ","))

}
