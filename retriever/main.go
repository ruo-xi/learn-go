package main

import (
	"fmt"
	"learn-go/retriever/mock"
	real2 "learn-go/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "cname",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
	//s.Post()
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Content: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//type assertion
	//realRetriever := r.(*real2.Retriever)
	//fmt.Println(realRetriever.TimeOut)
	if mockRetirever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetirever.Content)
	} else {
		fmt.Println("ERROR")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(">%T     %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(">Contents", v.Content)
	case *real2.Retriever:
		fmt.Println(">UserAgent:", v.UserAgent)
	}
	println()
}
