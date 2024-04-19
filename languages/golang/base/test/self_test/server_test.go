package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

const (
	SERVER_TEST          int = 0 // 測試
	SERVER_GRPC_INTERNAL int = 0 // grpc 内部
	SERVER_GRPC_EXTERNAL int = 0 // grpc 外部
	SERVER_HTTP          int = 0 // http
	runMode              int = SERVER_TEST
)

var tm Manager

func TestMain(m *testing.M) {
	fmt.Println("全局测试前执行操作,初始化服务、数据库及创建数据等")
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	setup()
	retCode := m.Run()
	teardown()
	fmt.Println("全局测试前执行操作,关闭服务、数据库及删除数据等")
	os.Exit(retCode)
}

// 开始及结束操作
func setup() {
	switch runMode {
	case SERVER_TEST:
		m := Manager{}
		m.Init()
		tm = m
	}

	logrus.Info("setup start")

}

func teardown() {
	logrus.Info("setup end")
}

func TestSendGetRequest(t *testing.T) {
	tm.Get("xiao hong")
}
