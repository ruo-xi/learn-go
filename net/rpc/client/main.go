package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"ruo-xi/learn-go/net/rpc/server/hello_service"
)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloServiceByDefault(network string, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		log.Fatal(err)
	}
	return &HelloServiceClient{Client: c}, err
}

func DialHelloServiceByJSON(network string, address string) (*HelloServiceClient, error) {
	netClient, err := net.Dial(network, address)
	if err != nil {
		log.Fatal(err)
	}
	c := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(netClient))
	return &HelloServiceClient{Client: c}, err
}

func (c HelloServiceClient) Hello(request string, reply *string) error {
	return c.Client.Call(hello_service.HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloServiceByJSON("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply string

	client.Call("HelloService.Hello", "hello", &reply)
	//err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

}

//func main() {
//	client, err := rpc.Dial("tcp", "localhost:1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//
//	var reply string
//	err = client.Call("HelloService.Hello", "hello", &reply)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(reply)
//}
