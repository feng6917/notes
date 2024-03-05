package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func main() {
	/*
		Go http
		与我们一般编写的http服务器不同，go为了实现高并发和高性能，使用了goroutines来处理conn的读写事件，这样每个请求都保持独立，相互不会阻塞，可以高效对的响应网络事件
	*/
	http.HandleFunc("/get", PrintGet)
	http.HandleFunc("/post-body", PrintPostBody)
	http.HandleFunc("/post-form", PrintPostForm)
	http.HandleFunc("/post-form-file", PrintPostFormFile)
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

func PrintPostBody(w http.ResponseWriter, r *http.Request) {
	var bs struct {
		Name string `json:"name"`
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &bs)
	if err != nil {
		panic(err)
	}
	fmt.Println("name: ", bs.Name)
	w.Write([]byte(fmt.Sprintf("{name:%s}", bs.Name)))
}

func SendPostBodyRequest(name string) (string, error) {
	/*
			 postValue := url.Values{
			      "emal":{"xx@xx.com"},
			   }
			postString := postValue.Encode()
		    strings.NewReader(postString)
	*/
	m := make(map[string]interface{}, 0)
	m["name"] = name
	mb, _ := json.Marshal(m)
	resp, err := http.Post("http://127.0.0.1:8080/post-body", "application/x-www-form-urlencoded", bytes.NewBuffer(mb))
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func PrintPostForm(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	name, nameOk := r.Form["name"]
	if !nameOk {
		fmt.Println("name not ok")
	}
	w.Write([]byte(fmt.Sprintf("{name:%s}", name)))
}

func SendPostFormRequest(name string) (string, error) {
	form := url.Values{}

	form.Add("name", name)
	resp, err := http.PostForm("http://127.0.0.1:8080/post-form", form)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func PrintPostFormFile(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	_, multiFileHeader, err := r.FormFile("f")
	if err != nil {
		panic(err)
	}
	fmt.Println(multiFileHeader.Filename)
	w.Write([]byte(fmt.Sprintf("{name:%s}", multiFileHeader.Filename)))
}

func SendPostFormFileRequest(fPath string) (string, error) {
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	f, err := os.Open(fPath)
	if err != nil {
		panic(err)
	}
	_, fname := filepath.Split(fPath)
	wfile, err := writer.CreateFormFile("f", fname)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(wfile, f)
	if err != nil {
		panic(err)
	}
	_ = writer.Close() // 发送之前必须调用Close()以写入结尾行
	request, err := http.NewRequest("POST", "http://127.0.0.1:8080/post-form-file", buffer)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{
		Timeout: time.Duration(10) * time.Minute, // 超时时间10分钟
	}
	resp, err := client.Do(request)
	if resp == nil {
		panic(err)
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(result), nil
}
