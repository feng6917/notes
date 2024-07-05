package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/duke-git/lancet/v2/random"
)

func main() {
	/*
		 	cond 当一组操作条件都满足时 开启继续执行接下来的操作 wait （wait 必须加锁 必须对所有条件全部满足进行判断） broadcast 每次满足条件进行广播 signal 允许调用者caller唤醒一个等待此Cond的goroutine

			模拟剧组拍戏各组准备 道具组、灯光组、导演组、美术组、摄影组、剧务组
	*/

	sc := &sync.Cond{L: &sync.Mutex{}}

	var ready int32

	for i := 0; i < 6; i++ {
		go func(i int) {
			time.Sleep(time.Duration(random.RandInt(1, 3)) * time.Second)
			switch i {
			case 0:
				fmt.Println("道具组 准备完毕！")
			case 1:
				fmt.Println("灯光组 准备完毕！")
			case 2:
				fmt.Println("导演组 准备完毕！")
			case 3:
				fmt.Println("美术组 准备完毕！")
			case 4:
				fmt.Println("摄影组 准备完毕！")
			case 5:
				fmt.Println("剧务组 准备完毕！")
			default:
				fmt.Println("制片组 准备完毕！")
			}
			sc.L.Lock()
			ready++
			sc.L.Unlock()
			sc.Broadcast()
		}(i)
	}

	sc.L.Lock()
	for ready != 6 {
		sc.Wait()
		fmt.Println("--- 导演耐心等待中 ---")
	}
	sc.L.Unlock()

	fmt.Println("========== action")
}
