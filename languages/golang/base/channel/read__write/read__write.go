package read__write

import (
	"fmt"
	"time"
)

// Ch1WriteAndRead 无缓存ch 数据写入后，需要进行读取，数据不取出来，无法再次写入
func Ch1WriteAndRead() {
	var ch = make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("writedata: ", i)
		}
		close(ch)
		fmt.Println("close chan!")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Millisecond)
			v, ok := <-ch
			if ok {
				fmt.Println("read data: ", v)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Exit!")
}

// Ch2WriteAndRead 有缓存ch 数据写入后，可以在缓存容量内再次写入，缓存区为空，读取阻塞
func Ch2WriteAndRead() {
	var ch = make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("writedata: ", i)
		}
		close(ch)
		fmt.Println("close chan!")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Millisecond)
			v, ok := <-ch
			if ok {
				fmt.Println("read data: ", v)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Exit!")
}

// Ch1CloseRead 读取已关闭无缓存无元素ch 返回类型零值 第二个返回值为false
func Ch1CloseRead() {
	var ch = make(chan int)
	close(ch)
	go func() {
		for i := 0; i < 5; i++ {
			val, ok := <-ch
			fmt.Printf("val: %v, success: %v\n", val, ok)
		}
	}()
	time.Sleep(time.Second)
}

// Ch2CloseRead 读取已关闭有缓存有元素ch 有元素返回元素值，第二个值返回true，元素读取完成后返回类型零值 第二个返回值为false
func Ch2CloseRead() {
	var ch = make(chan bool, 2)
	ch <- true
	close(ch)
	go func() {
		for i := 0; i < 5; i++ {
			val, ok := <-ch
			fmt.Printf("val: %v, success: %v\n", val, ok)
		}
	}()
	time.Sleep(time.Second)
}
