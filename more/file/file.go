package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {

	/*
		路径
			相对路径 relatice
				相对于当前工程
			绝对路径 absolute
	*/
	//路径
	fileName1 := "/home/yu/Desktop/learn-go/fib.txt"
	fileName2 := "fib.txt"
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	fmt.Println(filepath.Abs(fileName2))
	fmt.Println("获取父目录:", path.Join(fileName1, ".."))

	//创建目录
	createDir()

	//创建文件
	err, done := createFile()
	if done {
		return
	}
	//打开文件
	file3, done2 := openFile(err, fileName2)
	if done2 {
		return
	}

	//关闭文件
	file3.Close()

	//删除文件或文件夹
	removeFile(err)
}

func createDir() {
	//创建目录
	os.Mkdir("a", os.ModePerm)
	os.MkdirAll("a/b/c/d", os.ModePerm)
}

func createFile() (error, bool) {
	file1, err := os.Create("a/a1.txt")
	//采用模式0666创建一个名为name的文件，如果文件已存在会截断它（为空文件）。
	//如果成功，返回的文件对象可用于I/O；
	//对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。
	if err != nil {
		fmt.Println("err:", err)
		return nil, true
	}
	fmt.Println(file1)
	return err, false
}

func openFile(err error, fileName2 string) (*os.File, bool) {
	/*
		打开文件
			os.Open(filename)
			os.OpenFile(filename, mode, perm)
		filename/文件名
		mode/打开方式
			const (
				O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
				O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
				O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
				O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
				O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
				O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
				O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
				O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
			)
		perm/文件不存在时,创建文件指定权限
	*/

	file2, err := os.Open(fileName2)
	//只读
	if err != nil {
		fmt.Println("err", err)
		return nil, true
	}
	fmt.Println(file2)

	file3, err := os.OpenFile(fileName2, os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	//可选
	if err != nil {
		fmt.Println("err:", err)
		return nil, true
	}
	fmt.Println(file3)
	return file3, false
}

func removeFile(err error) {
	err = os.Remove("a/a1.txt")

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	err = os.Remove("a/b/c/d")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	err = os.Remove("a")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	err = os.RemoveAll("a")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
