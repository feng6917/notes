package task

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Live *sync.Map
}

func (c *Task) Init() {
	c.Live = &sync.Map{}
	go func() {
		tk := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-tk.C:
				c.PrintTask()
			}
		}
	}()
}

func (c *Task) Exist(key int64) bool {
	_, exist := c.Live.Load(key)
	return exist
}

func (c *Task) Ping(key int64, val string) {
	c.Live.Store(key, val)
}

func (c *Task) Remove(key int64) {
	c.Live.Delete(key)
}

func (c *Task) RangeTask() (int64, string, bool) {
	var resKey int64
	var resVal string
	var ok bool
	c.Live.Range(func(key, value interface{}) bool {
		keyVal, exist := key.(int64)
		if exist {
			if time.Unix(keyVal, 0).Sub(time.Now()) < 0 {
				resKey = keyVal
				resVal, _ = value.(string)
				ok = true
				return true
			}
		}
		return true
	})

	return resKey, resVal, ok
}

func (c *Task) PrintTask() {
	var index int32
	fmt.Println("=========================TASK==================================")
	c.Live.Range(func(key, value interface{}) bool {
		keyVal, exist := key.(int64)
		if exist {
			index += 1
			fmt.Printf("待执行任务 数量: %d 时间: %s \r\n", index, time.Unix(keyVal, 0).Format("2006-01-02 15:04:05"))
		}
		return true
	})

}
