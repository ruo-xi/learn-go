package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "ccname",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 = empty map

	var m3 map[string]int //m3 = nil

	fmt.Println(m, m2, m3)

	//key 是无序的
	//key 需要可以判断是否相等
	//除了slice map function
	//struct不包含上述字段时可以作为key 编译是检测是否包含
	for k, v := range m {
		fmt.Println(k, v)
	}
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	//取不存在的key value取zeroValue
	courseName, ok = m["course1"]
	fmt.Println(courseName, ok)
	fmt.Printf("%q", courseName)
	println()

	//综上
	if name, ok := m["course"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exist")
	}

	delete(m, "name")
	name, ok := m["name"]
	fmt.Println(name, ok)
}
