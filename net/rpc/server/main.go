package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"ruo-xi/learn-go/net/rpc/server/hello_service"
)

func main() {
	hello_service.RegisterHelloService(new(hello_service.HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		//defer conn.Close()
		fmt.Println(conn.RemoteAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
		//fmt.Println("over")
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
