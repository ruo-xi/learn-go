package main

import "fmt"

// slice是arr的一个视图
//type  struct {
//	ptr 指针 指向slice开头的元素
//	len  长度
//	cap array从ptr开始到结束的长度
//}
//slice可以向后扩展 不可以向前扩展
//

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[2:6] =", arr[:6])
	s1 := arr[2:]
	updateSlice(s1)
	fmt.Println("arr[2:6] =", s1)
	s2 := arr[:]
	updateSlice(s2)
	fmt.Println("arr[2:6] =", s2)
	fmt.Println(arr)

	arr[0], arr[2] = 0, 2
	fmt.Println(arr)
	s1 = arr[2:6]
	fmt.Println(s1)
	s2 = s1[3:5]
	fmt.Println(s1, len(s2), cap(s2))
	fmt.Println(s2[0:3])

	//超过cap长度后 会重新创建一个
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	//s4 s5 不再以arr为视图
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)

}

func updateSlice(s []int) {
	s[0] = 100
}
