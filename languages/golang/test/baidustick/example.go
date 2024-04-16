package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

/*
步骤
1. 获取cookie
2. 获取贴吧列表
3. 开始签到(获取tbs)
*/

func main() {
	// 1. 获取用户cookie
	fileBuf, err := os.ReadFile("./init.json")
	if err != nil {
		fmt.Println("读取配置文件失败,err: ", err)
		panic("--------------------------")
	}
	var cf struct {
		Cookie string
	}
	err = json.Unmarshal(fileBuf, &cf)
	if err != nil {
		fmt.Println("解析配置文件失败,err: ", err)
		panic("--------------------------")
	}
	fmt.Printf("成功解析到配置文件中Cookie, Cookie 长度: %d\r\n\r\n", len(cf.Cookie))

	// 3. 获取贴吧列表
	names, err := GetNameListRequest(cf.Cookie)
	if err != nil {
		fmt.Printf("ERROR: 获取用户贴吧列表失败, error:%v\r\n", err)
		panic("--------------------------")
	}
	fmt.Printf("成功通过请求获取到用户贴吧, 获取到数量: %d\r\n", len(names))

	fmt.Printf("\r\n\r\n开始签到\r\n\r\n")
	for index, tmp := range names {
		time.Sleep(time.Millisecond * 10)
		fmt.Printf("%d ----- %s 吧 开始签到\r\n", index, tmp)
		tbs, err := GetTBSRequest(cf.Cookie)
		if err != nil {
			fmt.Printf("ERROR: 获取用户贴吧TBS失败, error:%v\r\n", err)
			panic("--------------------------")
		}
		var hasSend bool
		err, hasSend = SendPostFormFileRequest(cf.Cookie, tmp, tbs)
		if err != nil {
			fmt.Printf("%s 吧 签到失败, err: %v\r\n", tmp, err)
			continue
		}
		if hasSend {
			fmt.Printf("%d ----- %s 吧 亲，你之前已经签过了\r\n\r\n\r\n", index, tmp)
		} else {
			fmt.Printf("%d ----- %s 吧 签到成功\r\n\r\n\r\n", index, tmp)
		}
	}
}

func GetNameListRequest(cookie string) ([]string, error) {
	var names []string
	request, err := http.NewRequest("GET", "http://tieba.baidu.com/f/like/mylike", nil)
	if err != nil {
		return names, err
	}
	request.Header.Set("Cookie", cookie)
	client := &http.Client{
		Timeout: time.Duration(1) * time.Minute, // 超时时间10分钟
	}
	resp, err := client.Do(request)
	if resp == nil {
		return names, err
	}

	body, err := DecodeHTMLBody(resp.Body)
	if err != nil {
		return names, err
	}
	// 加载 HTML document对象
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".forum_main .forum_table table tbody td").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		// title := s.Find("i").Text()
		if i%4 == 0 && band != "" {
			fmt.Printf("Review %d: %s \n", i, band)
			names = append(names, band)
		}
	})
	return names, nil
}

func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		if _, name, _ := charset.DetermineEncoding(data, ""); len(name) != 0 {
			return name
		}
	}

	return "utf-8"
}

func DecodeHTMLBody(body io.Reader) (io.Reader, error) {
	charset := detectContentCharset(body)

	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}

	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}

	return body, nil
}

func GetTBSRequest(cookie string) (string, error) {
	request, err := http.NewRequest("GET", "http://tieba.baidu.com/dc/common/tbs", nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Cookie", cookie)
	client := &http.Client{
		Timeout: time.Duration(1) * time.Minute, // 超时时间10分钟
	}
	resp, err := client.Do(request)
	if resp == nil {
		return "", err
	}

	defer resp.Body.Close()
	resBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var item struct {
		IsLogin int    `json:"is_login"`
		Tbs     string `json:"tbs"`
	}
	err = json.Unmarshal(resBuf, &item)
	if err != nil {
		return "", err
	}
	if item.IsLogin != 1 || item.Tbs == "" {
		fmt.Println("item: ", item)
		return "", errors.New("请求返回值有误")
	}

	return item.Tbs, nil
}

func SendPostFormFileRequest(cookie, kw, tbs string) (error, bool) {
	form := url.Values{}

	form.Add("ie", "utf-8")
	form.Add("kw", kw)
	form.Add("tbs", tbs)

	request, err := http.NewRequest("POST", "https://tieba.baidu.com/sign/add", bytes.NewBufferString(form.Encode()))
	if err != nil {
		return err, false
	}
	request.Header.Set("Cookie", cookie)
	client := &http.Client{
		Timeout: time.Duration(10) * time.Minute, // 超时时间10分钟
	}
	resp, err := client.Do(request)
	if resp == nil {
		return err, false
	}

	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, false
	}
	if strings.Contains(string(result), `"no":1101`) {

		return nil, true
	}
	var item struct {
		Error string
		Data  struct {
			Errmsg string `json:"errmsg"`
		}
	}

	err = json.Unmarshal(result, &item)
	if err != nil {
		return err, false
	}
	if item.Error != "" || item.Data.Errmsg != "success" {
		return err, false
	}
	return nil, false
}
