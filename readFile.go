package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("zzzzz")
	if err != nil {
		log.Fatal(err)
	}

	bytes := make([]byte, 20)
	file.Read(bytes)
	fmt.Println(bytes)
	bytes1 := make([]byte, 20)
	file.Read(bytes1)
	fmt.Println(bytes1)
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	println(fileInfo.Size())
}
