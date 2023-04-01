package main

import (
	"fmt"
	"time"

	broadcast "github.com/dustin/go-broadcast"
)

func main() {
	// 初始化
	b := broadcast.NewBroadcaster(100)
	r := NewDataEvent(b)

	// 发布（生产）数据
	r.Publish()

	// 订阅（消费）数据
	r.Subscribe()

	time.Sleep(11 * time.Second)
}

// 定义数据结构
type DataEvent struct {
	Broadcaster broadcast.Broadcaster
}

func NewDataEvent(bc broadcast.Broadcaster) *DataEvent {
	return &DataEvent{Broadcaster: bc}
}

func (c *DataEvent) Publish() {
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				tmp := time.Now().Format("2006-01-02 15:04:05")
				c.Broadcaster.Submit(tmp)
				fmt.Printf("publish data: %v \r\n", tmp)
			}
		}
	}()
}

func (c *DataEvent) Subscribe() {
	ch := make(chan interface{})
	go func() {
		c.Broadcaster.Register(ch)
		defer c.Broadcaster.Unregister(ch)
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("subscribe data: %v \r\n", <-ch)
			}
		}
	}()
}
