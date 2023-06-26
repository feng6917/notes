package util

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
  时间获取
  时间转换必须加入时区设置,请注意
  @author Bill
*/

var (
	randSeek = int64(1)
	l        sync.Mutex
	zone     = "CST" //时区
)

func TimeIntToDate(time_int int) string {
	var cstZone = time.FixedZone(zone, 8*3600)
	return time.Unix(int64(time_int), 0).In(cstZone).Format("2006-01-02 15:04:05")
}

func GetNowDateTime() string {
	var cstZone = time.FixedZone(zone, 8*3600)
	return time.Now().In(cstZone).Format("2006-01-02 15:04:05")
}

func GetDate() string {
	var cstZone = time.FixedZone(zone, 8*3600)
	return time.Now().In(cstZone).Format("2006-01-02")
}

// 防时间间隔
func GetIntTime() int {
	var _t = int(time.Now().Unix())
	return _t
}

// 暂时独立
func _getRandomSring(num int, str ...string) string {
	s := "123456789"
	if len(str) > 0 {
		s = str[0]
	}
	l := len(s)
	r := rand.New(rand.NewSource(getRandSeek()))
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		x := r.Intn(l)
		buf.WriteString(s[x : x+1])
	}
	return buf.String()
}

func getRandSeek() int64 {
	l.Lock()
	if randSeek >= 100000000 {
		randSeek = 1
	}
	randSeek++
	l.Unlock()
	return time.Now().UnixNano() + randSeek
}

// 获取今天时间戳 Today => 00:00:00
func TodayTimeUnix() int {
	t := time.Now()
	tm1 := int(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix())
	return tm1
}

// 获取今天时间戳 Today => 23:59:59
func TodayNightUnix() int {
	tm1 := TodayTimeUnix() + 86400 - 1
	return tm1
}

type UTime struct{}

// 现在时间
func (c *UTime) NowTime() time.Time {
	return time.Now()
}

// 现在时间戳 13位（纳秒数）
func (c *UTime) NowUnix() int64 {
	return time.Now().UnixNano() / 1e6
}

// 解析时间戳13位（纳秒数） 为时间
func (c *UTime) ParseUnixToTime(un int64) time.Time {
	return time.Unix(un/1000, 0)
}

// 时间戳差值比较 秒
func (c *UTime) CompareTimeUnix(sourceUnix, destUnix int64) float64 {
	t1 := c.ParseUnixToTime(sourceUnix)
	t2 := c.ParseUnixToTime(destUnix)
	d := t1.Sub(t2)
	return d.Seconds()
}

// 时间增加若干秒
func (c *UTime) AddTime(t time.Time, aYear, aMonth, aDay, aHour, aMinute, aSecond int) time.Time {
	if aYear != 0 || aMonth != 0 || aDay != 0 {
		t = t.AddDate(aYear, aMonth, aDay)
	} else if aHour != 0 {
		hd, _ := time.ParseDuration(fmt.Sprintf("%dh", aHour))
		t = t.Add(hd)
	} else if aMinute != 0 {
		md, _ := time.ParseDuration(fmt.Sprintf("%dm", aMinute))
		t = t.Add(md)
	} else if aSecond != 0 {
		sd, _ := time.ParseDuration(fmt.Sprintf("%ds", aSecond))
		t = t.Add(sd)
	}
	return t
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func (c *UTime) GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func (c *UTime) GetLastDateOfMonth(d time.Time) int {
	return c.GetFirstDateOfMonth(d).AddDate(0, 1, -1).Day()
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 时间戳(纳秒 13位)转指定格式时间
func (c *UTime) ToSetTime(un int64) string {
	ts := time.Unix(un/1000, 0).Format("2006-01-02 15:04:05")
	return ts
}
