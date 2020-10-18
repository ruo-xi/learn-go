package hello_service

import "net/rpc"

const HelloServiceName = "HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}
type HelloService struct {
}

func (receiver *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}
func RegisterHelloService(hsi HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, hsi)
}
