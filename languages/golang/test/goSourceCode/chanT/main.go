package main

import (
	"fmt"
	"time"
)

/*
// 有缓存 发送发生在接收之前
// 无缓存 接收发生在发送之前

// 读取Nil Chan 会一直阻塞，关闭Nil Chan 会Panic
// 关闭已经关闭的Chan 会Panic
// 向关闭的Chan 写入数据会Panic

*/

func main() {

	// ChanObstructWrite()

	// 阻塞chan,读取数据
	ChanObstructRead()
	time.Sleep(time.Hour * 24) //
	// ChanObstruct()
}

func ChanObstructWrite() {
	ch := make(chan int)
	ch <- 1

}

func ChanObstructRead() {
	// ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// }()
	// v := <-ch
	// fmt.Println(v)
	// v1 := <-ch
	// fmt.Println(v1)
	var ch chan int
	go func() {
		fmt.Println("===== Read Nil Chan Before =======")
		v := <-ch
		fmt.Println(v)
	}()

	close(ch)
}

func ChanObstruct() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
	}()
	v := <-ch
	v1 := <-ch
	fmt.Println(v, v1)
}

func ChanNoObstruct() {
	ch := make(chan int, 1)
	select {
	case v := <-ch:
		fmt.Println(v)
		break
	default:
		fmt.Println("No data")
	}

	go func() { ch <- 1 }()
}
