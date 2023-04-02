package sbreak

import (
	"fmt"
	"time"
)

func BreakA(ch chan int) {
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("val: ", v)
				if v == 2 {
					break
				}
			}
		}
	}
	fmt.Println("exit!")
}

// BreakA1 for select 无法直接使用break 关闭, 可以使用 break+label 关闭
func BreakA1(ch chan int) {
Exit:
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("val: ", v)
				if v == 2 {
					break Exit
				}
			}
		}
	}
	fmt.Println("exit!")
}

// BreakA2 for select 可以使用chan关闭
func BreakA2(ch chan int) {
	chStatus := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-ch:
				if ok {
					fmt.Println("val: ", v)
					if v == 2 {
						break
					}
				}
			// 使用time.After 仅为判断条件
			case <-time.After(10 * time.Millisecond):
				chStatus <- true
				close(chStatus)
				break
			}
		}
	}()
	<-chStatus
	fmt.Println("exit!")
}
