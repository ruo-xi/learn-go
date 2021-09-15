package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入:")
	for scanner.Scan() {
		if scanner.Text() == "q" {
			break
		}
		_, err := conn.Write([]byte(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		fmt.Println(n)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("received")
		fmt.Println("请输入:")
	}
}
