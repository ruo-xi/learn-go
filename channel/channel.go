package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	//var c chan int // n == nil
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	//for i := 0; i < 10; i++ {
	//	channels[i] = make(chan int)
	//	go worker(i, channels[i])
	//}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		//n := <-channels[i]
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
	//for true {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}

	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

// <-chan int
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a' + 0
	c <- 'a' + 1
	c <- 'a' + 2
	c <- 'a' + 3
	c <- 'a' + 4
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a' + 0
	c <- 'a' + 1
	c <- 'a' + 2
	c <- 'a' + 3
	c <- 'a' + 4
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	fmt.Println("Channel as first-class citizen")
	//chanDemo()
	fmt.Println("Buffered channel")
	//bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
