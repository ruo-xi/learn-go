package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Content: "this is a fake laonaailifa.top"}
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//type assertion
	//realRetriever := r.(*real2.Retriever)
	//fmt.Println(realRetriever.TimeOut)
	if mockRetirever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetirever.Content)
	} else {
		fmt.Println("ERROR")
	}
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents", v.Content)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
