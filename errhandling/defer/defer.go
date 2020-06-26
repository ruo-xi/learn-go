package main

import (
	"bufio"
	"fmt"
	"learn-go/functional/fib"
	"os"
)

//defer 栈 先进后出
//场景 lock/unlock   open/close printheader/printfooter
func tryDefer() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	////panic("error occurred")
	////return
	//fmt.Println(4)
	//defer fmt.Println(5)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("print too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//writeFile("errhandling/defer/fib.txt")

	tryDefer()
	//fmt.Sprintf("123")

}
