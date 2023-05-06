package broadcast

import (
	"github.com/dustin/go-broadcast"
	"testing"
	"time"
)

func TestNewDataEvent(t *testing.T) {
	// 初始化
	b := broadcast.NewBroadcaster(100)
	r := NewDataEvent(b)
	if r == nil {
		t.Error("nil pointer ")
	}
}

func TestDataEvent_Publish(t *testing.T) {
	// 初始化
	b := broadcast.NewBroadcaster(100)
	r := NewDataEvent(b)
	if r == nil {
		t.Error("nil pointer ")
	}

	// 发布（生产）数据
	var v int
	v = 0
	r.Publish(&v)

	time.Sleep(3 * time.Second)
	want := 3
	if v != want {
		t.Errorf("got %d want %d", v, want)
	}
}

func TestDataEvent_Subscribe(t *testing.T) {
	// 初始化
	b := broadcast.NewBroadcaster(100)
	r := NewDataEvent(b)
	if r == nil {
		t.Error("nil pointer ")
	}

	// 发布（生产）数据
	var i int
	r.Publish(&i)

	// 订阅（消费）数据
	var v int
	r.Subscribe(&v)

	time.Sleep(3 * time.Second)
	want := 1
	if v != want {
		t.Errorf("got %d want %d", v, want)
	}
}
