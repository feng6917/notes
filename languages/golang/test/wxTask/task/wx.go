package task

import (
	"errors"
	"fmt"
	"os"

	"github.com/eatmoreapple/openwechat"
	"github.com/sirupsen/logrus"
)

type WX struct {
	bot *openwechat.Bot
}

func (c *WX) Init() {
	bot := openwechat.DefaultBot()

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")

	defer reloadStorage.Close()

	// 执行热登录
	bot.HotLogin(reloadStorage)

	// 获取登陆的用户
	_, err := bot.GetCurrentUser()
	if err != nil {
		logrus.Errorf("GetCurrentUser failed, err: %v", err)
		return
	}
	c.bot = bot
	return
}

func (c *WX) Login(isDesktop bool) error {
	if c.bot != nil {
		_, err := c.bot.GetCurrentUser()
		if err == nil {
			err = errors.New("用户已经登录")
			logrus.Info(err)
			return err
		}
	}
	os.RemoveAll("storage.json")
	bot := openwechat.DefaultBot(openwechat.Normal)
	if isDesktop {
		bot = openwechat.DefaultBot(openwechat.Desktop)
	}

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	if err := bot.Login(); err != nil {
		logrus.Error("============= login failed, ", err)
		return err
	}
	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")

	defer reloadStorage.Close()

	// 执行热登录
	bot.HotLogin(reloadStorage)

	// 获取登陆的用户
	_, err := bot.GetCurrentUser()
	if err != nil {
		logrus.Errorf("GetCurrentUser failed, err: %v", err)
		return err
	}
	c.bot = bot
	return nil
}

func (c *WX) IsLogin() error {
	if c.bot == nil {
		return errors.New("bot is nil")
	}
	self, err := c.bot.GetCurrentUser()
	if err != nil {
		return err
	}
	if self == nil {
		return errors.New("self is nil")
	}
	return nil
}

func (c *WX) ExistRemarkName(name string) bool {
	self, err := c.bot.GetCurrentUser()
	if err != nil {
		return false
	}
	// 获取所有的好友
	friends, err := self.Friends()
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 搜索备注名称
	f := friends.SearchByRemarkName(1, name)
	return f.Count() > 0
}

func (c *WX) SendMsgByRemarkName(name, msg string) {
	self, err := c.bot.GetCurrentUser()
	if err != nil {
		return
	}
	// 获取所有的好友
	friends, err := self.Friends()
	if err != nil {
		logrus.Errorf("self friends failed, err: %v", err)
		return
	}

	// 搜索备注名称
	f := friends.SearchByRemarkName(1, name)
	if f.Count() > 0 {
		err = f.SendText(msg)
		if err != nil {
			logrus.Errorf("send msg failed, err: %v", err)
			return
		}
	}
}
