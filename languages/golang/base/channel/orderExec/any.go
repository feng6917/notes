package orderExec

import (
	"fmt"
	"time"
)

func any() {
	// 多个chan实现
	// 洗手
	var chHandWashing = make(chan struct{}, 1)
	// 做饭
	var chCook = make(chan struct{}, 1)
	// 吃饭
	var chDine = make(chan struct{}, 1)

	chHandWashing <- struct{}{}
	close(chHandWashing)

	go func() {
		_, ok := <-chDine
		if ok {
			fmt.Println("开始吃饭！")
			fmt.Println("END!")
		}
	}()

	go func() {
		_, ok := <-chCook
		if ok {
			fmt.Println("开始做饭，准备吃饭！")
			chDine <- struct{}{}
			close(chDine)
		}
	}()

	go func() {
		_, ok := <-chHandWashing
		if ok {
			fmt.Println("开始洗手，准备做饭！")
			chCook <- struct{}{}
			close(chCook)
		}
	}()

	time.Sleep(1 * time.Second)
}
