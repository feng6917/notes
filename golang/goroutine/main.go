package main

import (
	"context"
	"fmt"
	"goroutine/close"
	"goroutine/gtimeout"
	"sync"
	"time"
)

func main() {}

func mClose() {
	// close.CloseA
	// 1. 定义一个chan, 写入n个值, 写入完成关闭
	// 2. 并发读取，读取到特定值，读取结束 break
	// 需要单独判断值，不建议使用

	// close.CloseA1
	// 1. 并发读取，读取到特定值，读取结束 break
	// 2. 并发写入，写入完成，关闭
	// 写入阻塞，chan未关闭，不建议使用

	// close.CloseA2
	// 同A1,多了一个判断chan,当读取完成时，设置该chan为true,
	// 每次写入时获取判断chan值，为true停止写入，关闭chan
	close.CloseA2()

}

func mTimeout() {
	// 简单使用测试
	// gtimeout.GTimeoutDoCtxTime()

	/*
		场景：两个goroutine,第一个2秒执行完毕，第二个3秒执行完毕。
	*/

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		go gtimeout.GTimeoutF1(ch1)
		select {
		case <-ctx.Done():
			fmt.Println("f1 timeout")
			break
		case <-ch1:
			fmt.Println("f1 Done")
		}
	}()

	go func() {
		go gtimeout.GTimeoutF2(ch2)
		select {
		case <-ctx.Done():
			fmt.Println("f2 timeout")
			break
		case <-ch2:
			fmt.Println("f2 done")
		}
	}()

	time.Sleep(time.Second * 5)
}

func mGWaitgroup() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 2000; i > 0; i-- {
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()
}
