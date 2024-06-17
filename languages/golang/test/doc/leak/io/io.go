package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		go httpGet()
		w.Write([]byte("Hello, world!"))
	})
	fmt.Println("http server start")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func httpGet() {

	// resp, err := http.Get("http://google.com")
	// if err != nil {
	// 	// logrous.Error()
	// 	log.Println("err: ", err)
	// }
	// 设置请求时长
	cli := http.Client{
		Timeout: time.Second,
	}
	reqUrl := "http://google.com"
	httpReq, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	httpResp, err := cli.Do(httpReq)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer httpResp.Body.Close()
}
