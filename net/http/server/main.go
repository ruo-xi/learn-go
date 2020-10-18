package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sayhello)
	server := http.Server{
		Addr:    ":1234",
		Handler: mux,
	}
	fmt.Println(server.ListenAndServe())
	server.ListenAndServe()
}

func sayhello(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "hello")
}
