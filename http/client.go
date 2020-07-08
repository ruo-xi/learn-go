package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet,
		"https://www.imooc.com",
		nil)
	request.Header.Add()
	client := http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			//return nil // 重定向
			//retuen err //终止重定向
			fmt.Println("Redirect:", req)
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}
	resp, err := client.Do(request)
	//resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get("https://m.imooc.com")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)

}
