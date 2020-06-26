package main

import (
	"fmt"
	"learn-go/oop/queue"
)

func main() {
	q := queue.Queue{1, 2, 3}
	q.Push(4)
	q.Push(5)
	q.Push("abc")
	q.Push(6)
	for !q.IsEmpty() {
		fmt.Println(q)
		println(q.Pop())
	}

}
