- 新浪 定时请求数据测试
  - 参考链接：
  - https://www.jianshu.com/p/108b8110a98c
  - https://www.cnblogs.com/zeroes/p/sina_stock_api.html
-  操作方式：
   1. 连接dongdong数据库
   2. 在bases 表中添加sz,sh 代码
   3. 刷新msgs 数据库
- 效果图：
  <img width="1280" alt="image" src="https://github.com/feng6917/notes/assets/82997695/1e402b71-fb40-40c9-bb7c-6873616f9aa9">

- code
```
  package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Base struct {
	Code string `gorm:"primaryKey;comment:代码"` // 股票代码加沪深代码 sz002307,sh600928
}

/*
0：”大秦铁路”，股票名字；
1：”27.55″，今日开盘价；
2：”27.25″，昨日收盘价；
3：”26.91″，当前价格；
4：”27.55″，今日最高价；
5：”26.20″，今日最低价；
6：”26.91″，竞买价，即“买一”报价；
7：”26.92″，竞卖价，即“卖一”报价；
8：”22114263″，成交的股票数，由于股票交易以一百股为基本单位，所以在使用时，通常把该值除以一百；
9：”589824680″，成交金额，单位为“元”，为了一目了然，通常以“万元”为成交金额的单位，所以通常把该值除以一万；
10：”4695″，“买一”申请4695股，即47手；
11：”26.91″，“买一”报价；
12：”57590″，“买二”
13：”26.90″，“买二”
14：”14700″，“买三”
15：”26.89″，“买三”
16：”14300″，“买四”
17：”26.88″，“买四”
18：”15100″，“买五”
19：”26.87″，“买五”
20：”3100″，“卖一”申报3100股，即31手；
21：”26.92″，“卖一”报价
(22, 23), (24, 25), (26,27), (28, 29)分别为“卖二”至“卖四的情况”
30：”2008-01-11″，日期；
31：”15:05:32″，时间；
*/

type Msg struct {
	Code         string  `gorm:"primaryKey;comment:代码"` // 代码
	Name         string  // 名称
	StartPrice   float64 // 开盘价
	CurrentPrice float64 // 当前价
	Change       string  // 变动
	TopPrice     string  // 最高价
	UnderPrice   string  // 最低价
	Buy1         string  // 买一
	Buy2         string  // 买二
	Buy3         string  // 买三
	Buy4         string  // 买四
	Buy5         string  // 买五
	Sale1        string  // 卖一
	Sale2        string  // 卖二
	Sale3        string  // 卖三
	Sale4        string  // 卖四
	Sale5        string  // 卖五
	MinImage     string  // 分时图
	Time         string  // 时间
}

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("./dongdong.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&Base{}, &Msg{})
	_ = db.Model(&Base{}).Create(&Base{Code: "sz002307"}).Error

	DB = db
}

func main() {
	fmt.Println("dong dong")
	// 4. 定时请求
	go func() {
		t := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println("recover", err)
						}
					}()
				}()

				proc()
			}
		}
	}()

	select {}

}

func proc() {
	// 1. 获取数据库 要查询的股票代码
	res, err := getDbBase()
	if err != nil {
		return
	}
	if len(res) <= 0 {
		return
	}
	log.Infof("list msg req count: %d", len(res))
	// 2. 获取股票基本信息
	msgs, err := getBaseMsg(res)
	if err != nil {
		log.Errorf("get base msg fail, err: ", err)
		return
	}
	_ = DB.Model(&Msg{}).Where("code != ''").Delete(&Msg{}).Error
	// 3. 写入数据库
	err = DB.Model(&Msg{}).CreateInBatches(msgs, 100).Error
	if err != nil {
		log.Errorf("CreateInBatches msg fail, err: ", err)
		return
	}
}

func getDbBase() ([]string, error) {
	var res []string
	err := DB.Model(&Base{}).Select("code").Pluck("code", &res).Error
	if err != nil {
		fmt.Println("db get base fail, err: ", err)
		return nil, err
	}
	return res, nil
}

func getBaseMsg(req []string) ([]*Msg, error) {
	if len(req) <= 0 {
		return nil, nil
	}
	getUrl := "https://hq.sinajs.cn/list="
	for index, tmp := range req {
		if index != 0 {
			getUrl += ","
		}
		getUrl += tmp
	}
	request, err := http.NewRequest(http.MethodGet, getUrl, nil)
	if err != nil {
		log.Errorf("NewRequest fail, err: %v", err)
		return nil, err
	}
	request.Header.Set("Referer", "https://finance.sina.com.cn")
	client := &http.Client{
		Timeout: time.Duration(10) * time.Minute, // 超时时间10分钟
	}
	resp, err := client.Do(request)
	if resp == nil {
		log.Errorf("client.Do fail, err: %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if resp == nil {
		log.Errorf("ReadAll fail, err: %v", err)
		return nil, err
	}
	resStr := string(result)
	var resList []string
	if strings.Contains(resStr, ";") {
		resList = strings.Split(resStr, ";")
	} else {
		resList = append(resList, resStr)
	}
	var msgs []*Msg
	for _, tmp := range resList {
		tmpRes, err := splitBody(tmp)
		if err != nil {
			continue
		}
		if tmpRes == nil {
			continue
		}
		if tmpRes.Code == "" {
			continue
		}
		msgs = append(msgs, tmpRes)
	}
	return msgs, nil
}

func splitBody(req string) (*Msg, error) {
	res := &Msg{}
	if !strings.Contains(req, "var hq_str_") {
		return res, nil
	}
	req = strings.ReplaceAll(req, "var hq_str_", "")
	l1 := strings.Split(req, `="`)
	if len(l1) != 2 {
		return res, nil
	}
	res.Code = l1[0]
	l2 := strings.Split(l1[1], ",")
	for index, tmp := range l2 {
		if index == 0 {
			res.Name = tmp
		} else if index == 2 {
			pr, err := strconv.ParseFloat(tmp, 64)
			if err != nil {
				log.Errorf("StartPrice parse fail, err: %v", err)
			}
			res.StartPrice = pr
		} else if index == 3 {
			pr, err := strconv.ParseFloat(tmp, 64)
			if err != nil {
				log.Errorf("CurrentPrice parse fail, err: %v", err)
			}
			res.CurrentPrice = pr
		} else if index == 4 {
			res.TopPrice = tmp
		} else if index == 5 {
			res.UnderPrice = tmp
		} else if index == 10 {
			res.Buy1 = tmp
		} else if index == 11 {
			res.Buy1 = tmp + " - " + getTr(res.Buy1)
		} else if index == 12 {
			res.Buy2 = tmp
		} else if index == 13 {
			res.Buy2 = tmp + " - " + getTr(res.Buy2)
		} else if index == 14 {
			res.Buy3 = tmp
		} else if index == 15 {
			res.Buy3 = tmp + " - " + getTr(res.Buy3)
		} else if index == 16 {
			res.Buy4 = tmp
		} else if index == 17 {
			res.Buy4 = tmp + " - " + getTr(res.Buy4)
		} else if index == 18 {
			res.Buy5 = tmp
		} else if index == 19 {
			res.Buy5 = tmp + " - " + getTr(res.Buy5)
		} else if index == 20 {
			res.Sale1 = tmp
		} else if index == 21 {
			res.Sale1 = tmp + " - " + getTr(res.Sale1)
		} else if index == 22 {
			res.Sale2 = tmp
		} else if index == 23 {
			res.Sale2 = tmp + " - " + getTr(res.Sale2)
		} else if index == 24 {
			res.Sale3 = tmp
		} else if index == 25 {
			res.Sale3 = tmp + " - " + getTr(res.Sale3)
		} else if index == 26 {
			res.Sale4 = tmp
		} else if index == 27 {
			res.Sale4 = tmp + " - " + getTr(res.Sale4)
		} else if index == 28 {
			res.Sale5 = tmp
		} else if index == 29 {
			res.Sale5 = tmp + " - " + getTr(res.Sale5)
		} else if index == 30 {
			res.Time = tmp
		} else if index == 31 {
			res.Time = res.Time + " " + tmp
		}
	}
	// change := res.CurrentPrice - res.StartPrice
	var change string
	pr := ((res.CurrentPrice - res.StartPrice) / res.StartPrice) * 100
	// if pr >= 0 {
	// 	change += "+"
	// } else {
	// 	change += "-"
	// }
	res.Change = change + fmt.Sprintf("%.2f", pr)
	res.Name, _ = GbkToUtf8(res.Name)
	res.MinImage = "https://image.sinajs.cn/newchart/min/n/" + res.Code + ".gif"
	return res, nil
}

func GbkToUtf8(s string) (string, error) {
	reader := transform.NewReader(bytes.NewBufferString(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d), nil
}

func getTr(s string) string {
	ti, _ := strconv.Atoi(s)
	tr := ti / 100
	return strconv.Itoa(tr)
}


```
