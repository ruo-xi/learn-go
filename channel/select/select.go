package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 chan int // c1 and c2 == nil
	c1, c2 = generator(), generator()
	w := createWorker(0)
	//n := 0
	var values []int
	//hasValue := false
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var acticeValue int
		//if hasValue {
		//	activeWorker = w
		//}
		if len(values) > 0 {
			activeWorker = w
			acticeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
			//hasValue = true
			//fmt.Println("Received from c1:", n)
		case n := <-c2:
			values = append(values, n)
			//hasValue = true
		case activeWorker <- acticeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len =:", len(values))
		case <-tm:
			fmt.Println("Bye")
			return
			//fmt.Println("Received from c2:", n)
			//default:
			//	fmt.Println("No value received")
			//}
		}
	}
}
