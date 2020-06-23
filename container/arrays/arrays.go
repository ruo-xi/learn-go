package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var gird [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(gird)

	//len(arr3)
	//for i, v := range arr3 {
	//	fmt.Println(i, v)
	//}

	fmt.Println(arr1)
	printArray(arr1)
	fmt.Println(arr1)
	//printArray(arr2)   WRONG
}

//数组是值类型
//[10]int [20]int是不同的类型
//
//调用func printArray(arr [5]int)
//会拷贝数组
func printArray(arr [5]int) {

	fmt.Println(arr)
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
	fmt.Println(arr)
}

//func main() {
//	arrays()
//}
