package broadcast

import (
	"fmt"
	"time"

	broadcast "github.com/dustin/go-broadcast"
)

// DataEvent 定义数据结构
type DataEvent struct {
	Broadcaster broadcast.Broadcaster
}

func NewDataEvent(bc broadcast.Broadcaster) *DataEvent {
	return &DataEvent{Broadcaster: bc}
}

func (c *DataEvent) Publish(v *int) {
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				tmp := time.Now().Format("2006-01-02 15:04:05")
				c.Broadcaster.Submit(tmp)
				*v += 1
				fmt.Printf("publish data: %v \r\n", tmp)
			}
		}
	}()
}

func (c *DataEvent) Subscribe(v *int) {
	ch := make(chan interface{})
	go func() {
		c.Broadcaster.Register(ch)
		defer c.Broadcaster.Unregister(ch)
		ticker := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-ticker.C:
				*v += 1
				fmt.Printf("subscribe data: %v \r\n", <-ch)
			}
		}
	}()
}
