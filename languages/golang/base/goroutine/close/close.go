package close

import (
	"fmt"
	"time"
)

func CloseA() {
	var ch = make(chan int)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value == 1 {
					fmt.Println("Exit!")
					break
				}
				fmt.Println(value)
			}
		}
	}()

	ch <- 0
	ch <- 1
	close(ch)

	time.Sleep(1 * time.Second)
}

func CloseA1() {
	var ch = make(chan int)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value%9 == 0 {
					fmt.Println("Exit!")
					break
				}
				fmt.Println(value)
			}
		}
	}()

	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			t := <-ch
			fmt.Println("---- other: ", t)
		}
	}()

	time.Sleep(8 * time.Millisecond)
}

func CloseA2() {
	var ch = make(chan int)
	var chS = make(chan bool)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value%9 == 0 {
					fmt.Println("Exit!")
					chS <- true
					close(chS)
					break
				}
				fmt.Println(value)
			}
		}
	}()

	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		value := <-chS
		if value {
			close(ch)
		}
	}()

	time.Sleep(1 * time.Second)
}

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
	CloseA2()

}
