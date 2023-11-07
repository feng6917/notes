- 贴吧自动签到测试
- 操作逻辑：
  - 使用本人登录Cookie，获取到用户贴吧列表，遍历列表进行签到。
- 操作步骤
  -  获取cookie
  -  请求协议获取用户id
  -  请求协议获取用户关注贴吧列表
  -  遍历贴吧列表 签到（需要获取tbs）
 
- 配置文件：init.json
    ```
      {
        "Cookie": "BAIDUID=945E9886C9A948095ED63CF6ACBDCEC0:FG=1; cn=https%3A%2F%2Ffclog.baidu.com%2Flog%2Fweirwood%3Ftype%3Dperf\""
      }
    ```
- 代码：
    ```
        package main

        import (
        	"bytes"
        	"encoding/json"
        	"errors"
        	"fmt"
        	"io"
        	"net/http"
        	"net/url"
        	"os"
        	"time"
        )
        
        /*
        步骤
        1. 获取cookie
        2. 获取用户id
        3. 获取贴吧列表
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
        	fmt.Printf("成功解析到配置文件中Cookie, Cookie 长度: %d\r\n", len(cf.Cookie))
        	// 2. 获取用户id
        	userID, err := GetUserIDRequest(cf.Cookie)
        	if err != nil {
        		fmt.Printf("ERROR: 获取用户id失败, err:%v\r\n", err)
        		panic("--------------------------")
        	}
        	fmt.Printf("成功通过请求获取到用户ID: %d\r\n", userID)
        	// 3. 获取贴吧列表
        	names, err := GetNameListRequest(cf.Cookie, userID)
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
        		err = SendPostFormFileRequest(cf.Cookie, tmp, tbs)
        		if err != nil {
        			fmt.Printf("%s 吧 签到失败, err: %v\r\n", tmp, err)
        		}
        		fmt.Printf("%d ----- %s 吧 签到成功\r\n\r\n\r\n", index, tmp)
        	}
        }
        
        func GetUserIDRequest(cookie string) (int, error) {
        	request, err := http.NewRequest("GET", "https://tieba.baidu.com/mo/q/sync", nil)
        	if err != nil {
        		return 0, err
        	}
        	request.Header.Set("Cookie", cookie)
        	client := &http.Client{
        		Timeout: time.Duration(1) * time.Minute, // 超时时间10分钟
        	}
        	resp, err := client.Do(request)
        	if resp == nil {
        		return 0, err
        	}
        
        	defer resp.Body.Close()
        	resBuf, err := io.ReadAll(resp.Body)
        	if err != nil {
        		return 0, err
        	}
        
        	var us struct {
        		Error string
        		Data  struct {
        			UserID int `json:"user_id"`
        		}
        	}
        	err = json.Unmarshal(resBuf, &us)
        	if err != nil {
        		return 0, err
        	}
        	if us.Error != "success" {
        		return 0, err
        	}
        	return us.Data.UserID, nil
        }
        
        func GetNameListRequest(cookie string, userID int) ([]string, error) {
        	var names []string
        	request, err := http.NewRequest("GET", fmt.Sprintf("https://tieba.baidu.com/p/getLikeForum?uid=%d", userID), nil)
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
        
        	defer resp.Body.Close()
        	resBuf, err := io.ReadAll(resp.Body)
        	if err != nil {
        		return names, err
        	}
        
        	type Info struct {
        		ForumName string `json:"forum_name"`
        	}
        
        	var item struct {
        		ErrMsg string
        		Data   struct {
        			Info []Info `json:"info"`
        		}
        	}
        	err = json.Unmarshal(resBuf, &item)
        	if err != nil {
        		return names, err
        	}
        	if item.ErrMsg != "success" {
        		return names, err
        	}
        	if len(item.Data.Info) > 0 {
        		for _, tmp := range item.Data.Info {
        			if tmp.ForumName != "" {
        				names = append(names, tmp.ForumName)
        			} else {
        				continue
        			}
        		}
        	}
        	return names, nil
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
        
        func SendPostFormFileRequest(cookie, kw, tbs string) error {
        	form := url.Values{}
        
        	form.Add("ie", "utf-8")
        	form.Add("kw", kw)
        	form.Add("tbs", tbs)
        
        	request, err := http.NewRequest("POST", "https://tieba.baidu.com/sign/add", bytes.NewBufferString(form.Encode()))
        	if err != nil {
        		return err
        	}
        	request.Header.Set("Cookie", cookie)
        	client := &http.Client{
        		Timeout: time.Duration(10) * time.Minute, // 超时时间10分钟
        	}
        	resp, err := client.Do(request)
        	if resp == nil {
        		return err
        	}
        
        	defer resp.Body.Close()
        	result, err := io.ReadAll(resp.Body)
        	if err != nil {
        		return err
        	}
        	var item struct {
        		Error string
        		Data  struct {
        			Errmsg string `json:"errmsg"`
        		}
        	}
        
        	err = json.Unmarshal(result, &item)
        	if err != nil {
        		return err
        	}
        	if item.Error != "" || item.Data.Errmsg != "success" {
        		return err
        	}
        	return nil
        }

    ``` 
