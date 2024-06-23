package main

import (
	"context"
	"fmt"
	"time"
	"unicode/utf8"
	"unsafe"

	"github.com/sirupsen/logrus"
)

func main() {
	// test 变量声明
	// testVar()
	// test range
	// testRange()
	// testRangeSlice()
	// 测试数组
	// testArrayEdit()
	// // 测试切片
	// testSliceEdit()

	// // 测试字符串数量
	// testStringCount()

	// // 测试for退出
	// testForExit()
	// // 测试未初始化chan插入读取
	// testNotInitChanRW()

	// // 测试结构体值修改
	// testStructEdit()

	// // 测试for值异常
	// testForValErr()

	// // test map noInit
	// // testMapNoInit()

	// // test map init
	// testMapInit()

	// // test defer exec
	// testDeferExec()

	// 测试make 与 new 初始化map区别
	// testMakeNew()

	// 单引号包括的是rune类型 双引号包括的是字符串类型
	// 汉字占三个字节
	// testStrCopy()

	// testStructEditVale()
	// 多路复用 超时处理
	// testSelect()

	//
	// testContextDead()
	// testContextTimeOut()
	// testCtxWithValue()
	// testChan()
	// testRwMtx()

	// testReverseStr()

	testPerson()

}

type Person interface {
	GetAge()
}

type Student struct {
	Name string
	Age  int
}

func (c *Student) GetAge() {
	fmt.Println("Student Age: ", c.Age)
}

type Techer struct {
	Hobby string
	Age   int
}

func (c *Techer) GetAge() {
	fmt.Println("Techer Age: ", c.Age)
}

func testPerson() {
	p := Person(&Student{"wangYi", 18})
	p.GetAge()

	t := Person(&Techer{"play", 24})
	t.GetAge()
}

func testReverseStr() {
	s := "aaa1123测试sf"
	var ns string
	for i := range []rune(s) {
		tmp := string([]rune(s)[len([]rune(s))-i-1])
		fmt.Println(tmp)
		ns += tmp
	}
	fmt.Println(ns)
}

func testChan() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

	// select {
	// case msg := <-ch:
	// 	fmt.Println(msg)
	// default:
	// 	fmt.Println("default")
	// }

	// select {}
}

func testRwMtx() {

	/*

		w           Mutex        // held if there are pending writers //
		writerSem   uint32       // semaphore for writers to wait for completing readers // 等待完成写 排队的信号量
		readerSem   uint32       // semaphore for readers to wait for completing writers // 等待完成读 排队的信号量
		readerCount atomic.Int32 // number of pending readers // 读锁的计数器 2 30 次方 最大数量
		readerWait  atomic.Int32 // number of departing readers // 等待读锁释放数量 逐渐递减为0
	*/
	// rw := sync.RWMutex{}

	// go func() {
	// 	// read
	// 	rw.RLock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("read 1 0: ", t1, t2, t3, t4)
	// 	fmt.Println("read 1 0")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 1 1")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 1 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("read 1: ", t1, t2, t3, t4)
	// 	rw.RUnlock()

	// }()
	// time.Sleep(time.Second)
	// go func() {
	// 	rw.RLock()
	// 	fmt.Println("read 2 0")
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("read 2 0: ", t1, t2, t3, t4)
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 2 1")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 2 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("read 2 2: ", t1, t2, t3, t4)
	// 	rw.RUnlock()
	// }()
	// time.Sleep(time.Millisecond)
	// go func() {
	// 	rw.RLock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("read 3 0: ", t1, t2, t3, t4)
	// 	fmt.Println("read 3 0")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 3 1")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 3 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("read 3 2: ", t1, t2, t3, t4)
	// 	rw.RUnlock()
	// }()
	// go func() {
	// 	rw.Lock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("write 1 0: ", t1, t2, t3, t4)
	// 	fmt.Println("write 1 0")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("write 1 1")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("write 1 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("write 1 2: ", t1, t2, t3, t4)
	// 	rw.Unlock()
	// }()
	// time.Sleep(time.Second * 3)
	// go func() {
	// 	rw.Lock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("write 2 0: ", t1, t2, t3, t4)
	// 	fmt.Println("write 2 0")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("write 2 1")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("write 2 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("write 2 2: ", t1, t2, t3, t4)
	// 	rw.Unlock()
	// }()
	// time.Sleep(time.Second)
	// go func() {
	// 	rw.RLock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("read 4 0: ", t1, t2, t3, t4)
	// 	fmt.Println("read 4 0")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 4 1")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 4 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("read 4 2: ", t1, t2, t3, t4)
	// 	rw.RUnlock()
	// }()
	// time.Sleep(time.Millisecond)
	// go func() {
	// 	rw.RLock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("read 5 0: ", t1, t2, t3, t4)
	// 	fmt.Println("read 5 0")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 5 1")
	// 	time.Sleep(time.Second)
	// 	fmt.Println("read 5 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("read 5 2: ", t1, t2, t3, t4)
	// 	rw.RUnlock()
	// }()
	// go func() {
	// 	rw.RLock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("read 6 0: ", t1, t2, t3, t4)
	// 	fmt.Println("read 6 0")
	// 	fmt.Println("read 6 1")
	// 	fmt.Println("read 6 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("read 6 2: ", t1, t2, t3, t4)
	// 	rw.RUnlock()
	// }()
	// go func() {
	// 	rw.RLock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("w 0: ", t1, t2, t3, t4)
	// 	fmt.Println("w 0")
	// 	fmt.Println("w 1")
	// 	fmt.Println("w 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("w 2: ", t1, t2, t3, t4)
	// 	rw.RUnlock()
	// }()
	// go func() {
	// 	rw.RLock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("w1 0: ", t1, t2, t3, t4)
	// 	fmt.Println("w1 0")
	// 	fmt.Println("w1 1")
	// 	fmt.Println("w1 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("w1 2: ", t1, t2, t3, t4)
	// 	rw.RUnlock()
	// }()
	// go func() {
	// 	rw.RLock()
	// 	t1, t2, t3, t4 := rw.GetInfo()
	// 	fmt.Println("w2 0: ", t1, t2, t3, t4)
	// 	fmt.Println("w2 0")
	// 	fmt.Println("w2 1")
	// 	fmt.Println("w2 2")
	// 	t1, t2, t3, t4 = rw.GetInfo()
	// 	fmt.Println("w2 2: ", t1, t2, t3, t4)
	// 	rw.RUnlock()
	// }()

	select {}
}

var neverChan = make(chan struct{})

func testContextDead() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	select {
	case <-neverChan:
		fmt.Println("receive msg from neverChan")
	case <-ctx.Done():
		fmt.Println("ctx done. ", "ctx err: ", ctx.Err())
	}
}

func testContextTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	select {
	case <-neverChan:
		fmt.Println("receive msg from neverChan")
	case msg := <-ctx.Done():
		fmt.Println("msg: ", msg)
		fmt.Println("ctx err: ", ctx.Err())
	}
}

func testCtxWithValue() {
	ctx := context.Background()
	nctx := context.WithValue(ctx, "name", "zhangsan")

	f := func(key string) {
		valI := nctx.Value(key)
		if valI == nil {
			return
		}
		valTmp, ok := valI.(string)
		fmt.Println(valTmp, ok)
	}

	f("name")
}

func testSelect() {
	// 多路复用
	ts1()
	// 超时处理
	ts2()
}

func ts2() {
	ch := make(chan int)
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case msg := <-time.After(time.Second):
		fmt.Println("time: ", msg)
	}
}

func ts1() {
	ch := make(chan int)
	go func() {
		ch <- 0
	}()

	go func() {
		ch <- 1
	}()

	select {
	case msg := <-ch:
		fmt.Println("msg: ", msg)
	case msg1 := <-ch:
		fmt.Println("msg1: ", msg1)
	}
}

type TS struct {
	Name string
}

func testStructEditVale() {
	ts := TS{Name: "aaa"}
	fmt.Printf("addr: %p\r\n", &ts)
	ts.Name = "bbb"
	ts.Name = "safdasfsfsadfffasfdsafafsfd"
	fmt.Printf("addr: %p\r\n", &ts)

}

func testStrCopy() {
	s1 := "aaaaa"
	sb1 := []byte(s1)
	s1n := string(sb1)
	fmt.Printf("s1:%#v sb1:%#v s1n:%#v\r\n", &s1, &sb1, &s1n)

	s2 := "bbbbbb"
	sr2 := []rune(s2)
	s2n := string(sr2)
	fmt.Printf("s2: %#v sr2:%#v s2n:%#v\r\n", &s2, &sr2, &s2n)

	s3 := 'a'
	fmt.Println(s3)
	fmt.Println(unsafe.Sizeof(s3))

	s4 := "ha哈哈"
	s4r := []rune(s4)
	s4b := []byte(s4)
	fmt.Println("rune: ", s4r)
	fmt.Println("byre: ", s4b)
}

func testMakeNew() {
	n := *new(map[string]int)
	n = map[string]int{}
	n["1"] = 1
	fmt.Println("n: ", n)

	m := make(map[string]int)
	m["1"] = 1
	fmt.Println("m: ", m)
}

func testDeferExec() {
	i := [2]int{1, 2}
	fmt.Println("before i: ", i)
	defer func() {
		fmt.Println("defer print i: ", i)
	}()
	i[0] = 2
	fmt.Println("after i: ", i)
}

func testMapInit() {
	m := make(map[string]int, 1)
	m["1"] = 1
	fmt.Println(m)
	m["2"] = 2
	fmt.Println(m)
}

func testMapNoInit() {
	var m map[string]int
	m["2"] = 1 // err
	fmt.Println(m)
}

func testRange() {
	lm := make(map[int]string)
	lm[1] = "1"
	lm[2] = "2"
	lm[3] = "3"
	lm[4] = "4"
	for k, v := range lm {
		fmt.Printf("k: %d k addr: %X v:%s v addr: %X \n", k, &k, v, &v)
	}
}

func testRangeSlice() {
	/*
		range 时kv地址始终是不变的，元素的副本，引用地址发生变化
	*/
	l := []int{1, 2, 3, 4}
	for k, v := range l {
		fmt.Printf("k: %d k addr: %X v:%d v addr: %X v Addr: %X \n", k, &k, v, &v, &l[k])
	}
}

func testVar() {
	var a int
	fmt.Println(a)
	// var a, c int // err
	// fmt.Println(a, c)
	s := 1
	fmt.Println(s)
	s, sn := 2, 3
	fmt.Println(s, sn)
}

func testForValErr() {
	tsList := []TS{{"1"}, {"2"}, {Name: "3"}}
	for _, v := range tsList {
		// source
		go func() {
			v.print()
		}()
		// ans1
		// vcopy := v
		// go func() {
		// 	vcopy.print()
		// }()
		// ans2
		// go func(req TS) {
		// 	req.print()
		// }(v)
	}
	time.Sleep(time.Second)
}

func (c *TS) EditName(name string) {
	c.Name = name
}

func (c *TS) print() {
	fmt.Println(c.Name)
}

func testStructEdit() {
	ts := TS{Name: "1"}
	logrus.Info("edit before: ", ts)
	ts.EditName("hahah")
	logrus.Info("edit after: ", ts)
	time.Sleep(time.Second)
}

func testNotInitChanRW() {
	// var ch chan int
	// // ch <- 1
	// <-ch
	// time.Sleep(time.Second * 2)
}

func testForExit() {
	ch := make(chan struct{}, 1)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(i)
		if i == 99 {
			ch <- struct{}{}
		}
	}
	<-ch
}

func testStringCount() {
	s := "23k中国"
	logrus.Info("字符串字节数量 : ", len(s))
	logrus.Info("字符串 rune数量: ", utf8.RuneCountInString(s))

	for _, v := range []byte(s) {
		fmt.Printf("v: %#x\n", v)
	}

}

func testArrayEdit() {
	s := [5]int{1, 2, 3}
	logrus.Info("edit before: ", s)
	arrayEdit(s)
	logrus.Info("edit after: ", s)
}

func arrayEdit(s [5]int) {
	s[0] = 9
}

func testSliceEdit() {
	s := []int{1, 2, 3}
	logrus.Info("edit before: ", s)
	sliceEdit(s)
	logrus.Info("edit after: ", s)
}

func sliceEdit(s []int) {
	s[0] = 9
}
