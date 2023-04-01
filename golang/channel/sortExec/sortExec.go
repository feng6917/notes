package sortExec

import (
	"fmt"
	"time"
)

var ch = make(chan int, 1)

func SortExecA() {
	// 多个chan实现
	var chXS = make(chan struct{}, 1)
	var chZF = make(chan struct{}, 1)
	var chCF = make(chan struct{}, 1)

	chXS <- struct{}{}
	close(chXS)

	go func() {
		_, ok := <-chCF
		if ok {
			fmt.Println("开始做饭，准备吃饭！")
			fmt.Println("END!")
		}
	}()

	go func() {
		_, ok := <-chZF
		if ok {
			fmt.Println("开始做饭，准备吃饭！")
			chCF <- struct{}{}
			close(chCF)
		}
	}()

	go func() {
		_, ok := <-chXS
		if ok {
			fmt.Println("开始洗手，准备做饭！")
			chZF <- struct{}{}
			close(chZF)
		}
	}()

	time.Sleep(1 * time.Second)
}
