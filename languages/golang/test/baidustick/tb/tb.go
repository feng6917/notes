package tb

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
)

type Tb struct {
	Name       string
	CurExp     string
	BadgeTitle string
	BadgeLevel string
}

func GetTb(cookie string) ([]Tb, error) {
	var res []Tb
	for i := 1; i < 100; i++ {
		tmp, err := GetNameListRequest(cookie, i)
		if err != nil {
			return nil, err
		}
		if len(tmp) > 0 {
			res = append(res, tmp...)
		}
		if len(tmp) == 0 {
			break
		}
	}
	return res, nil
}

func GetNameListRequest(cookie string, pn int) ([]Tb, error) {
	var res []Tb
	request, err := http.NewRequest("GET", fmt.Sprintf("http://tieba.baidu.com/f/like/mylike?&pn=%d", pn), nil)
	if err != nil {
		return res, err
	}
	request.Header.Set("Cookie", cookie)
	client := &http.Client{
		Timeout: time.Duration(1) * time.Minute, // 超时时间10分钟
	}
	resp, err := client.Do(request)
	if resp == nil {
		return res, err
	}

	body, err := DecodeHTMLBody(resp.Body)
	if err != nil {
		return res, err
	}
	// 加载 HTML document对象
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		logrus.Errorf("err %v", err)
		// log.Fatal(err)
		return nil, err
	}

	// Find the review items
	var name, curExp, badgeTitle, badgeLevel string
	doc.Find(".forum_main .forum_table table tbody td").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if i%4 == 0 && s.Find("a").Text() != "" {
			name = s.Find("a").Text()
		}
		// title := s.Find("i").Text()
		if s.Find(".cur_exp").Text() != "" {
			curExp = s.Find(".cur_exp").Text()
		}
		if s.Find(".like_badge_title").Text() != "" {
			badgeTitle = s.Find(".like_badge_title").Text()
		}

		if s.Find(".like_badge_lv").Text() != "" {
			badgeLevel = s.Find(".like_badge_lv").Text()
		}
		if name != "" && badgeLevel != "" {
			fmt.Printf("Review %d: %s %s %s %s\n", i, name, curExp, badgeTitle, badgeLevel)
			res = append(res, Tb{
				Name:       name,
				CurExp:     curExp,
				BadgeTitle: badgeTitle,
				BadgeLevel: badgeLevel,
			})
			name = ""
			curExp = ""
			badgeTitle = ""
			badgeLevel = ""
		}
	})
	return res, nil
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
	fmt.Println("Rs: ", string(resBuf))
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
		return "", errors.New("isLogin failed, 请求返回值有误")
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
