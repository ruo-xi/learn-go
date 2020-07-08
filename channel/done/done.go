package main

import (
	"fmt"
	"sync"
)

func chanDemo() {
	//var c chan int // n == nil
	var workers [10]worker

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-worker.done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-worker.done
	}

	wg.Wait()

	//wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}

//func doWorker(id int, c chan int, done chan bool)
func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		go func() { w.done() }()
	}
}

type worker struct {
	in   chan int
	done func()
	//done chan bool
	//wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
		//done: make(chan bool),
		//wg: wg,
	}
	go doWorker(id, w)
	return w
}

func main() {
	chanDemo()
}
