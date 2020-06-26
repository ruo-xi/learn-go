package main

import "learn-go/functional/fib"

func main() {
	f := fib.Fibonacci()
	//for i := 0; i < 10; i++ {
	//	println(f())
	//}
	fib.PrintFileContents(f)

}
