package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	t := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 3,
			KeepAlive: time.Second * 3,
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       time.Second * 90, //空闲超时时间
		TLSHandshakeTimeout:   time.Second * 10, // tls握手超时时间
		ExpectContinueTimeout: time.Second * 1,  //100-continue状态吗超时时间
	}

	client := http.Client{
		Timeout:   time.Second * 30,
		Transport: t,
	}
	resp, err := client.Get("http://127.0.0.1:1234/hello")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer resp.Body.Close()
	bds, err := ioutil.ReadAll(resp.Body)
	fmt.Println("gg")
	fmt.Println(string(bds))
}
