package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/*
		读取数据
			Reader接口
			type Reader interface {
				Read(p []byte) (n int, err error)
			}
	*/
	//打开文件
	fileName := "fib.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//关闭文件
	defer file.Close()
	//读取文件
	bs := make([]byte, 4, 4)
	fmt.Println(bs)
	n, err := file.Read(bs)
	fmt.Println(bs)
	//fmt.Println(string(bs))
	fmt.Println(n)
	fmt.Println(err)

	ioutil.ReadAll()
}
