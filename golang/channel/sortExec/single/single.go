package single

import (
	"fmt"
	"time"
)

var a = make(chan int, 1)

func Single() {
	a <- 1
	go ChXS()
	go CHZF()
	go CHCF()

	time.Sleep(3 * time.Second)

}

func CHCF() {
	v := <-a
	if v == 3 {
		fmt.Println("吃饭，结束")
		close(a)
	} else {
		a <- v
	}

}

func ChXS() {
	v := <-a
	if v == 1 {
		fmt.Println("洗手，做饭")
		a <- 2
	} else {
		a <- v
	}
}

func CHZF() {
	v := <-a
	if v == 2 {
		fmt.Println("做饭，吃饭")
		a <- 3
	} else {
		a <- v
	}
}
