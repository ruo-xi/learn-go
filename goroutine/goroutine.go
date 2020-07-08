package main

import (
	"fmt"
	"time"
)

func printHello() {

}

func main() {
	var a [1000]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
