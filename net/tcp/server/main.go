package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		go HandleConn(conn)
	}

}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr())
	bs := []byte
	for {
		n, err := conn.Read(bs)
		if err != nil {
			fmt.Println(conn.RemoteAddr(), " --- err:", err)
			return
		}
		conn.Write(bs)
		fmt.Println(conn.RemoteAddr(), " --- bytes: ", n)
		fmt.Println(string(bs[:n]))
	}

}
