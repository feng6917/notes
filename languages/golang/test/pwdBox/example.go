package main

import (
	"fmt"
	"lgo/util/hex"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
)

const (
	splitStr string = "\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n" +
		"\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n" +
		"\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n" +
		"\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n" +
		"\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n" +
		"\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n" +
		"\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n" +
		"\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n" +
		"\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n\r\n"

	createPwdSql = `
CREATE TABLE
IF
	NOT EXISTS 'pwds' (
		"url_address" VARCHAR ( 256 ) NULL, -- 网址
		"login_name" VARCHAR ( 256 ) NULL, -- 名称
		"pwd" VARCHAR ( 256 ) NULL -- 密码
	)`

	saltValue string = "zxcavtbnwmlbkjhgfdesabnf"
)

type PWD struct {
	UrlAddress string
	LoginName  string
	Pwd        string
}

// 密码箱 获取到密码只有一条路, 输入网址->输入账号->获取密码
// 添加密码流程 请输出网址(网址不能为空)->请输入账号(账号不能为空)->请输入密码(密码不能为空)
// 获取密码流程 --->
func main() {

	db, err := gorm.Open(sqlite.Open("./pwd.db"), &gorm.Config{
		Logger: glog.Default.LogMode(glog.Silent),
	})
	if err != nil {
		fmt.Println("初始化数据库失败, 错误: ", err)
		return
	}

	err = db.Exec(createPwdSql).Error
	if err != nil {
		fmt.Println("初始化数据表失败, 错误: ", err)
		return
	}

ROOTLABEL:
	loginNameCount, urlAddressCount, err := allPwds(db)
	if err != nil {
		fmt.Println("统计数据表数据失败, 错误: ", err)
		return
	}

	s1 := getStr(strconv.Itoa(loginNameCount), " ")
	s2 := getStr(strconv.Itoa(urlAddressCount), " ")

	tip := "==========================================================\r\n" +
		"=                        密码箱                          =\r\n" +
		"= 存储网址数量: " + s2 + "=\r\n" +
		"= 存储账号数量: " + s1 + "=\r\n" +
		"=--------------------------------------------------------=\r\n" +
		"= 提示:                                                  =\r\n" +
		"= 1. 退出系统: 在任意输入位置输入 'Q/q' .                =\r\n" +
		"= 2. 重置系统: 在输入网址及输入登录名位置输入 '0'.       =\r\n" +
		"==========================================================\r\n"
	fmt.Println(tip)

	var urlAddress, loginName, pwd string
	fmt.Println("请输入网址:")
	fmt.Scan(&urlAddress)
	urlAddress = strings.TrimSpace(urlAddress)
	if urlAddress == "q" || urlAddress == "Q" {
		return
	}
	if urlAddress == "" {
		printErrorStr("网址不能为空")
		goto ROOTLABEL
	}
	if urlAddress != "0" && len(urlAddress) < 5 {
		printErrorStr("网址长度不能小于 '5' ")
		goto ROOTLABEL
	}

LOGINNAMETLABEL:
	fmt.Println("请输入登录名: ")
	fmt.Scan(&loginName)
	loginName = strings.TrimSpace(loginName)
	if loginName == "q" || loginName == "Q" {
		return
	}
	if loginName == "" {
		printErrorStr("登录名不能为空")
		goto LOGINNAMETLABEL
	}
	if loginName != "0" && len(loginName) < 3 {
		printErrorStr("登录名长度不能小于 '3'")
		goto LOGINNAMETLABEL
	}

	if urlAddress == "0" && loginName == "0" {
		var resetPwd string
		fmt.Println("密码箱即将重置,请输入指令 '1' 确认:")
		fmt.Scan(&resetPwd)
		if resetPwd == "q" || resetPwd == "Q" {
			return
		}
		if resetPwd == "1" {
			fmt.Println("\r\n接收到确认指令,开始删除.")
			fmt.Println("\r\n网址查询中 .....")
			time.Sleep(time.Second * 1)
			fmt.Println("\r\n网址已删除")

			fmt.Println("\r\n登录名查询中 .....")
			time.Sleep(time.Second * 1)
			fmt.Println("\r\n登录名已删除")

			fmt.Println("\r\n密码查询中 .....")
			time.Sleep(time.Second * 1)
			fmt.Println("\r\n密码已删除")
			err = pwdDelete(db)
			if err != nil {
				printErrorStr("密码重置失败")
			}
			fmt.Println("\r\n密码箱重置成功.")
			gotoRoot()
			fmt.Println(splitStr)
		}
		goto ROOTLABEL
	}

	pwdInfo, err := pwdDetails(db, urlAddress, loginName)
	if err == nil && pwdInfo.Pwd != "" {
		bufPwd := []byte(pwdInfo.Pwd)
		var cbcBuf []byte
		cbcBuf, err = hex.AesDecryptCBC(bufPwd, []byte(saltValue))
		if err != nil {
			printErrorStr("解密失败")
			return
		}
		fmt.Println("\r\n****************************************")
		fmt.Println(string(cbcBuf))
		fmt.Println("****************************************\r\n")
		gotoRoot()
		goto ROOTLABEL
	}

PASSWORDLABEL:
	fmt.Println("请输入密码: ")
	fmt.Scan(&pwd)
	pwd = strings.TrimSpace(pwd)
	if pwd == "q" || pwd == "Q" {
		return
	}
	if pwd == "" {
		printErrorStr("密码不能为空")
		goto PASSWORDLABEL
	}
	if len(pwd) < 6 {
		printErrorStr("密码长度不能小于 '6' ")
		goto PASSWORDLABEL
	}

	bufPwd := []byte(pwd)
	var cbcBuf []byte
	cbcBuf, err = hex.AesEncryptCBC(bufPwd, []byte(saltValue))
	if err != nil {
		printErrorStr("加密失败")
		return
	}

	// 保存数据
	insertData := PWD{
		UrlAddress: urlAddress,
		LoginName:  loginName,
		Pwd:        string(cbcBuf),
	}
	if err := pwdInsert(db, &insertData); err != nil {
		printErrorStr("存入数据失败")
	}
	fmt.Println("存入数据成功")
	goto ROOTLABEL

}

func printErrorStr(errStr string) {
	str := "\r\nERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR ERROR\r\n"
	data := "ERROR     " + errStr
	fmt.Println(str + data + str)
	gotoRoot()
}

func allPwds(db *gorm.DB) (int, int, error) {
	var loginNameCount, urlAddressCount int64
	err := db.Table("pwds").Count(&loginNameCount).Error
	if err != nil {
		return 0, 0, err
	}
	err = db.Table("pwds").Group("url_address").Count(&urlAddressCount).Error
	return int(loginNameCount), int(urlAddressCount), err
}

func pwdDetails(db *gorm.DB, urlAddress, loginName string) (cd *PWD, err error) {
	err = db.Table("pwds").Where("url_address = ? and login_name = ?", urlAddress, loginName).First(&cd).Error
	return cd, err
}

func pwdInsert(db *gorm.DB, cr *PWD) error {
	err := db.Table("pwds").Create(&cr).Error
	return err
}

func pwdDelete(db *gorm.DB) error {
	err := db.Table("pwds").Where("url_address != ''").Delete(&PWD{}).Error
	return err
}

func gotoRoot() {
	fmt.Println("自动跳转到首页 ...\r\n       倒计时")
	time.Sleep(time.Second * 1)
	fmt.Println("        叁")
	time.Sleep(time.Second * 1)
	fmt.Println("        贰")
	time.Sleep(time.Second * 1)
	fmt.Println("        壹")
	fmt.Println(splitStr)
}

func getStr(s string, ns string) string {
	sl := len([]rune(s))
	l := 41 - sl
	if l > 0 {
		for i := 0; i < l; i++ {
			s += ns
		}
	}
	return s
}
