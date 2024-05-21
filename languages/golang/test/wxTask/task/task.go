package task

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Live *sync.Map
}

type TaskInfo struct {
	RemarkName string    // 备注
	Text       string    // 内容
	Time       time.Time // 发送开始时间
	Since      int       // 发送间隔
	Number     int       // 发送次数
	SendStatus string    // 发送状态
	SendNumber int       // 发送次数
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

func (c *Task) Exist(key string) bool {
	_, exist := c.Live.Load(key)
	return exist
}

func (c *Task) Ping(val TaskInfo) {
	c.Live.Store(val.RemarkName, val)
}

func (c *Task) Remove(key string) {
	c.Live.Delete(key)
}

func (c *Task) RangeTask() (TaskInfo, bool) {
	var res TaskInfo
	var ok bool
	c.Live.Range(func(key, val interface{}) bool {
		ti, exist := val.(TaskInfo)
		if exist {
			// 起始时间 + （间隔时间*分钟*待发送次数+1）
			if ti.SendNumber == ti.Number {
				fmt.Println("??????")
				c.Remove(ti.RemarkName)
			}
			ts := ti.Time.Add(time.Second * time.Duration(ti.Since) * (time.Duration(ti.SendNumber + 1)))
			if ts.Sub(time.Now()) < 0 {
				res = ti
				ok = true
				return true
			}
		}
		return true
	})

	return res, ok
}

func (c *Task) PrintTask() {
	var index int32
	fmt.Println("=========================TASK==================================")
	c.Live.Range(func(key, value interface{}) bool {
		task, exist := value.(TaskInfo)
		if exist {
			index += 1
			fmt.Printf("待执行任务 数量: %d 任务信息: %+v \r\n", index, task)
		}
		return true
	})

}
