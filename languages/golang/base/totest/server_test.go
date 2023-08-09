package main

import (
	"fmt"
	"os"
	"testing"
)

const (
	// repo-internal 测试方法直接启动repo服务进行测试
	// server-internal 测试方法直接启动server服务进行测试
	// server-auto 测试方法自动处理数据 模拟测试
	// server-grpc 调用已部署服务grpc测试
	// server-http 调用已部署服务http测试
	runMode string = ""
)

func TestMain(m *testing.M) {
	fmt.Println("全局测试前执行操作,初始化服务、数据库及创建数据等")
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()
	fmt.Println("全局测试前执行操作,关闭服务、数据库及删除数据等")
	os.Exit(retCode)
}

// 开始及结束操作
func setupTest(t *testing.T) func(t *testing.T) {
	t.Log("setup start")
	return func(t *testing.T) {
		t.Log("setup end")
	}
}


