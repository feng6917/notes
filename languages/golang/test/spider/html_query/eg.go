package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

func main() {
	demo1()
	// github author test
	demo2()
}

func demo1() {
	reqUrl := "http://tieba.baidu.com/f/like/mylike?&pn=1"
	request, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		logrus.Errorf("err %v", err)
		return
	}
	// request.Header.Set("Cookie", cookie)
	client := &http.Client{
		Timeout: time.Duration(1) * time.Minute, // 超时时间10分钟
	}
	resp, err := client.Do(request)
	if resp == nil {
		logrus.Errorf("err %v", err)
		return
	}

	body, err := DecodeHTMLBody(resp.Body)
	if err != nil {
		logrus.Errorf("err %v", err)
		return
	}
	// 加载 HTML document对象
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		logrus.Errorf("err %v", err)
		return
	}

	// Find the review items
	doc.Find(".forum_main .forum_table table tbody td").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if i%4 == 0 && s.Find("a").Text() != "" {
			fmt.Println(s.Find("a").Text())
		}
	})
}

func demo2() {
	reqUrl := "https://blog.csdn.net/ppdouble/article/details/134516917"
	request, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		logrus.Errorf("err %v", err)
		return
	}

	// 创建一个自定义的代理
	proxyURL, err := url.Parse("http://127.0.0.1:11809")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return
	}

	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// request.Header.Set("Cookie", cookie)
	client := &http.Client{
		Timeout:   time.Duration(1) * time.Minute, // 超时时间10分钟
		Transport: tr,
	}
	resp, err := client.Do(request)
	if resp == nil {
		logrus.Errorf("err %v", err)
		return
	}

	body, err := DecodeHTMLBody(resp.Body)
	if err != nil {
		logrus.Errorf("err %v", err)
		return
	}
	// 加载 HTML document对象
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		logrus.Errorf("err %v", err)
		return
	}

	// fmt.Println(doc.Text())

	// Find the review items
	s := doc.Find("#content_views > pre:nth-child(6) > code > span:nth-child(26)").Text()
	fmt.Println("s: ", s)
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
