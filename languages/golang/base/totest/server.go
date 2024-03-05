package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	t := Manager{}
	t.Run()
}

type Manager struct{
	s Service
}

func (c *Manager) Run() {
	runHttp()
}

func runHttp() {
	/*
		Go http
		与我们一般编写的http服务器不同，go为了实现高并发和高性能，使用了goroutines来处理conn的读写事件，这样每个请求都保持独立，相互不会阻塞，可以高效对的响应网络事件
	*/
	http.HandleFunc("/get", PrintGet)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func PrintGet(w http.ResponseWriter, r *http.Request) {
	// todo: 最原始的测试方法 待优化方法
	name := r.URL.Query().Get("name")
	fmt.Println("name: ", name)
	w.Write([]byte(fmt.Sprintf("{name:%s}", name)))
}

func SendGetRequest(name string) (string, error) {
	resp, err := http.Get("http://127.0.0.1:8080/get?name=" + name)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
